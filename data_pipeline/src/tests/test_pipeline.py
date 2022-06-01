from unittest import TestCase
from unittest.mock import patch

from pipeline.load import Loader
from pipeline.pipeline import run_pipeline
from pipeline.transform import Transformer


class TestPipeline(TestCase):
    def setUp(self) -> None:
        extract_data_patcher = patch('pipeline.extract.extract_data')
        transformer_patcher = patch.object(Transformer, 'get_new_cases_per_month_n_location')
        loader_patcher = patch.object(Loader, 'publish_record_list_to_db')
        self.addCleanup(extract_data_patcher.stop)
        self.addCleanup(transformer_patcher.stop)
        self.addCleanup(loader_patcher.stop)
        self.mock_extract = extract_data_patcher.start()
        self.mock_transformer = transformer_patcher.start()
        self.mock_loader = loader_patcher.start()

    def test_pipeline(self):
        run_pipeline()

        self.mock_transformer.assert_called_once()
        self.mock_loader.assert_called_once()
