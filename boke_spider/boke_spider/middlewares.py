from selenium import webdriver
from scrapy.http.response.html import HtmlResponse

class BokeSpiderDownloaderMiddleware:
    def __init__(self):
        self.driver = webdriver.Chrome(executable_path="D:\chromedriver\chromedriver.exe")

    # 拦截请求
    def process_request(self, request, spider):
        # 用selenium去请求
        self.driver.get(request.url)
        # todo 隐式等待
        self.driver.implicitly_wait(10)

        # 把selenium获得的网页对象，创建一个Response对象给spider
        response = HtmlResponse(request.url, body=self.driver.page_source, request=request, encoding='utf-8')
        return response
