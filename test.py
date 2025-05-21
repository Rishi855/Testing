import requests
import random
import string


def generate_random_code(prefix, total_length=5):
    characters = string.ascii_letters + string.digits
    suffix_length = total_length - len(prefix)
    suffix = ''.join(random.choices(characters, k=suffix_length))
    return prefix + suffix

def is_real_link(html):
    # Detect common interstitial marker phrases
    fake_markers = [
        "Antiphishing.biz", "not a robot", "cyber security", "landing page address will be shown",
        "confirm that you are not a robot"
    ]
    return not any(marker.lower() in html.lower() for marker in fake_markers)

base_url = "https://short-link.me/"
working_links = []

for _ in range(25):  # You can increase the number if needed
    code = generate_random_code('12')
    full_url = base_url + code
    try:
        response = requests.get(full_url, timeout=5)
        if response.status_code == 200 and is_real_link(response.text):
            print(f"✅ Valid: {full_url}")
            working_links.append(full_url)
        else:
            print(f"⚠️ Skipped: {full_url}")
    except Exception as e:
        print(f"❌ Error: {full_url} - {e}")

print("\n=== ✅ WORKING SHORT LINKS ===")
for link in working_links:
    print(link)
