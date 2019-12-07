import sys
import os
import pprint
pprint.pprint(sys.path)

sd = os.path.dirname(__file__)
sys.path.append(sd + '/../')
sys.path.append(sd + '/../../')

pprint.pprint(sys.path)
