import sys
import unittest

# sys.argv[1] = "python/y2023/d06/test_solution.py"
module = sys.argv[1].replace("/", ".").removesuffix(".py")
print(f"Running: {module}")

loader = unittest.TestLoader()
suite = loader.loadTestsFromName(module)
runner = unittest.TextTestRunner(verbosity=2)
runner.run(suite)
