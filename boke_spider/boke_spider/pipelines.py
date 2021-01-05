import pymysql
import time
from scrapy.utils.project import get_project_settings

from boke_spider.items import TagItem, BokeSpiderItem


class BokeSpiderPipeline:
    def __init__(self):
        settings = get_project_settings()
        host = settings['MYSQL_CONFIG']['HOST']
        port = settings['MYSQL_CONFIG']['PORT']
        user = settings['MYSQL_CONFIG']['USER']
        password = settings['MYSQL_CONFIG']['PASSWORD']
        database = settings['MYSQL_CONFIG']['DATABASE']
        self.db = pymysql.connect(host=host, port=port, user=user, password=password, database=database)

    def process_item(self, item, spider):
        if type(item) == BokeSpiderItem:
            now = time.strftime("%Y-%m-%d %H:%M:%S", time.localtime())
            cursor = self.db.cursor()
            # 查询
            searchSql = "select * from album_tag where album_id = %s"
            searchResult = cursor.execute(searchSql, item['albumId'])
            if searchResult:
                updateSql = "update album_tag set first_tag_id=%s,second_tag_id=%s,description=%s,updated_at=%s where album_id=%s"
                cursor.execute(updateSql,
                               (item['firstTagId'], item['secondTagId'], item['description'], now, item['albumId']))
                self.db.commit()
            else:
                insertSql = "insert into album_tag(id,first_tag_id,second_tag_id,album_id,description,created_at,updated_at) values(null,%s,%s,%s,%s,%s,%s)"
                cursor.execute(insertSql,
                               (item['firstTagId'], item['secondTagId'], item['albumId'], item['description'], now, now))
                self.db.commit()
        return item


class TagPipeline:
    def __init__(self):

        settings = get_project_settings()
        host = settings['MYSQL_CONFIG']['HOST']
        port = settings['MYSQL_CONFIG']['PORT']
        user = settings['MYSQL_CONFIG']['USER']
        password = settings['MYSQL_CONFIG']['PASSWORD']
        database = settings['MYSQL_CONFIG']['DATABASE']
        self.db = pymysql.connect(host=host, port=port, user=user, password=password, database=database)

    def process_item(self, item, spider):
        if type(item) == TagItem:
            now = time.strftime("%Y-%m-%d %H:%M:%S", time.localtime())
            cursor = self.db.cursor()
            searchSql = "select * from podcast_tag where id = %s"
            searchResult = cursor.execute(searchSql, item['tagId'])
            if searchResult:
                updateSql = "update podcast_tag set tag_name=%s,updated_at=%s where id=%s"
                cursor.execute(updateSql, (item['tagName'], now, item['tagId']))
                self.db.commit()
            else:
                insertSql = "insert into podcast_tag(id,tag_name,created_at,updated_at) values(%s,%s,%s,%s)"
                cursor.execute(insertSql, (item['tagId'], item['tagName'], now, now))
                self.db.commit()
        return item
