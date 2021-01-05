import scrapy


class BokeSpiderItem(scrapy.Item):
    firstTagId = scrapy.Field()
    secondTagId = scrapy.Field()
    albumId = scrapy.Field()
    description = scrapy.Field()


class TagItem(scrapy.Item):
    tagId = scrapy.Field()
    tagName = scrapy.Field()
