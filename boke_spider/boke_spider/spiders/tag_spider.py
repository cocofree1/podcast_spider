import scrapy
from boke_spider.items import BokeSpiderItem, TagItem
import re
from selenium import webdriver
from selenium.webdriver.chrome.options import Options
from selenium.webdriver import ChromeOptions
from scrapy.http.response.html import HtmlResponse
from scrapy.utils.project import get_project_settings

class TagSpiderSpider(scrapy.Spider):
    name = 'tag_spider'
    allowed_domains = ['https://podcasts.apple.com']
    start_urls = ['https://podcasts.apple.com/cn/genre/%E6%92%AD%E5%AE%A2/id26']

    def parse(self, response):
        chrome_options = Options()
        chrome_options.add_argument('--headless')
        chrome_options.add_argument('--disable-gpu')

        option = ChromeOptions()
        option.add_experimental_option('excludeSwitches', ['enable-automation'])
        driver = webdriver.Chrome(executable_path="D:\chromedriver\chromedriver.exe", chrome_options=chrome_options,
                                  options=option)
        url = "https://podcasts.apple.com/cn/genre/%E6%92%AD%E5%AE%A2/id26"
        driver.get(url)
        driver.implicitly_wait(10)
        tagResponse = HtmlResponse(url, body=driver.page_source, encoding='utf-8')

        # 获取所有标签
        tags = tagResponse.xpath("//div[@class='grid3-column']//a")
        for tag in tags:
            tagName = tag.xpath("./text()").get()
            tagUrl = tag.xpath("./@href").get()
            tagId = self.getIdByUrl(tagUrl)
            item = TagItem(tagId=tagId, tagName=tagName)
            yield item

        # 根据url获取专辑与标签的对应
        subjects = response.xpath("//div[@class='grid3-column']//@href").getall()
        # yield scrapy.Request(subjects[4], callback=self.parse_album, dont_filter=True)
        for subject in subjects:
            yield scrapy.Request(subject, callback=self.parse_tag, dont_filter=True)

    def parse_tag(self, response):
        # 获取一级,二级分类id,没有默认为0
        tags = response.xpath("//ul[@class='list breadcrumb']//@href").getall()
        firstTagId = self.getIdByUrl(tags[1])
        secondTagId = 0
        if len(tags) == 3:
            secondTagId = self.getIdByUrl(tags[2])

        # 获取专辑id
        albums = response.xpath("//div[@class='grid3-column']//@href").getall()
        for album in albums:
            albumId = self.getIdByUrl(album)
            yield scrapy.Request(album, callback=self.parse_album,
                                 meta={'firstTagId': firstTagId, 'secondTagId': secondTagId, 'albumId': albumId},
                                 dont_filter=True)

    def parse_album(self, response):
        description = response.xpath("//section[@class='product-hero-desc__section']//p/text()").get()
        description = re.sub(r'\n', '', str(description))
        description = re.sub(r' ', '', description)
        item = BokeSpiderItem(firstTagId=response.meta['firstTagId'],
                              secondTagId=response.meta['secondTagId'],
                              albumId=response.meta['albumId'],
                              description=description)
        yield item

    def getIdByUrl(self, url):
        parts = url.split("/")
        id_str = parts[len(parts) - 1]
        pattern = re.compile(r'\d+')
        ids = re.findall(pattern, id_str)
        id = int(ids[0])
        return id
