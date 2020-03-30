import re
import json
import os

from pythainlp import word_tokenize
from tqdm import tqdm

def arrange_lines(lines=[]):
    tail_pattern = r'.*,(\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}),(\d+),(\w+),(\d+),(.*)'
    cleaned_lines = []
    i = 0
    
    while i < len(lines):
        cleaned_line = ""
        m = re.match(tail_pattern, lines[i])
        
        while m is None and i < len(lines):
            cleaned_line += lines[i].replace('\n', '')
            m = re.match(tail_pattern, lines[i])
            i += 1

        i += 1
        cleaned_lines.append(cleaned_line)

        percentage = (((i) / len(lines)) * 100)
        print("Arranging progress %.4f%%" % percentage)
    
    return cleaned_lines

def to_json(lines=[]):
    dicts = []
    pattern = ','.join([
        r'^(?P<id>.*)',
        r'(?P<type>\w+)',
        r'(?P<message>.*)',
        r'(?P<time>\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2})',
        r'(?P<engagement>\d+)',
        r'(?P<channel>\w+)',
        r'(?P<owner_id>\w+)',
        r'(?P<owner_name>.*)$',
    ])

    for line in tqdm(lines):
        m = re.match(pattern, line)
        
        if m is None: continue
        
        groupdict = m.groupdict()
        groupdict['words'] = word_tokenize(groupdict['message'], keep_whitespace=False)
        groupdict['hashtags'] = re.findall(r'(#.*) ', groupdict['message'])

        dicts.append(groupdict)
    
    return json.dumps(dicts)


if __name__ == '__main__':
    with open('./rawdata.csv') as f:
        head = f.readline()
        lines = f.readlines()
    
    with open('./data.json', 'w') as f:
        data = to_json(arrange_lines(lines))
        f.write(data)