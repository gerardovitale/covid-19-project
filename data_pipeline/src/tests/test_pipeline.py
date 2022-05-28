from unittest import TestCase

from pipeline.pipeline import run_pipeline


class TestPipeline(TestCase):

    def test_pipeline(self):
        run_pipeline()
        self.assertTrue(True)
