from unittest import TestCase

from pipeline.transform import Transformer
from tests.resources import mock_data


class TestTransformer(TestCase):

    def test_get_new_cases_per_month_n_location_goes_well(self):
        test_data = mock_data('test-data.csv')
        expected_cols = ['year', 'month', 'location', 'total_cases', 'new_cases']
        expected_record = {'year': 2020, 'month': 2, 'location': 'Italy', 'total_cases': 3.0, 'new_cases': 0.0}

        actual_data = Transformer.get_new_cases_per_month_n_location(test_data)
        actual_cols = list(actual_data[0].keys())

        self.assertIsInstance(actual_data, list)
        self.assertIsInstance(actual_data[0], dict)
        self.assertEqual(expected_cols, actual_cols)
        self.assertEqual(expected_record, actual_data[0])
