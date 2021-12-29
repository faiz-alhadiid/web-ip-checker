import os
import re

def get_dig_result(url: str) -> str:
    stream = os.popen(f"dig {url}")
    return stream.read()

def parse_dig_result(text: str) -> list:
    rc = re.compile(r";; ANSWER SECTION:(.*?);;", re.MULTILINE| re.DOTALL)
    subs = rc.findall(text)

    if (len(subs) == 0):
        return [];

    answer_sections: str = subs[0].strip()
    records = answer_sections.split("\n")

    ips = []
    for record in records:
        row = record.split()
        ips.append({"type": row[-2],"value":row[-1]})
    return ips