from unittest import TestCase
from pandas import DataFrame

from tests.resources import EXPECTED_FEATURES
from pipeline.extract import extract_data


class TestExtract(TestCase):

    def test_extract_goes_well(self):
        data = extract_data(chunk=True)

        acutal_features = data.columns

        self.assertIsInstance(data, DataFrame)
        self.assertTrue((EXPECTED_FEATURES == acutal_features).all())
