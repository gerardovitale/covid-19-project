import os

import pymongo


class DataBaseConfig:

    def __init__(self, database='covid-19-project') -> None:
        self.client = pymongo.MongoClient(os.getenv("MONGO_PASS"))
        self.db = self.client.database
        self.collection_list = self.db.list_collection_names()
