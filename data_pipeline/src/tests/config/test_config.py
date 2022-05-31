from unittest import TestCase

from pymongo import MongoClient
from pymongo.database import Database

from pipeline.config.config import DataBaseConfig


class TestDataBaseConfig(TestCase):

    def test_object_instantiation_goes_well(self):
        mongo_db = DataBaseConfig()

        self.assertIsInstance(mongo_db.client, MongoClient)
        self.assertIsInstance(mongo_db.db, Database)
