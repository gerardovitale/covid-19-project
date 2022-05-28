from unittest import TestCase

from pandas import DataFrame

from pipeline.extract import extract_data
from tests.resources import EXPECTED_FEATURES


class TestExtract(TestCase):

    def test_extract_goes_well(self):
        expected_rows = 1000

        data = extract_data(chunk=True)

        self.assertIsInstance(data, DataFrame)
        self.assertTrue((EXPECTED_FEATURES == data.columns).all())
        self.assertEqual(expected_rows, data.shape[0])

    def test_extract_when_chunk_false(self):
        expected_rows = 178468
        expected_cols = 67

        data = extract_data(chunk=False)

        self.assertIsInstance(data, DataFrame)
        self.assertTrue((EXPECTED_FEATURES == data.columns).all())
        self.assertTrue(expected_rows <= data.shape[0])
        self.assertEqual(expected_cols, data.shape[1])
