from unittest import TestCase
from unittest.mock import patch
from unittest.mock import Mock

from pipeline.pipeline import run_pipeline


class TestPipeline(TestCase):

    @patch('pipeline.pipeline.run_pipeline')
    def test_pipeline(self, pipeline_mock: Mock):
        run_pipeline()
        self.assertTrue(True)
