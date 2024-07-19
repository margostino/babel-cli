import json
import os

import weaviate
from dotenv import load_dotenv
from weaviate.classes.query import MetadataQuery

load_dotenv()

OPENAI_API_KEY = os.environ.get("OPENAI_API_KEY")
METADATA_PATH = "/Users/martin.dagostino/.babel/db/metadata"

client = weaviate.connect_to_local()
# client = weaviate.Client(
#     url="http://localhost:8080",
#     additional_headers={"X-OpenAI-Api-Key": OPENAI_API_KEY},
# )

client.is_ready()


def get_all_json_files():
    json_list = []

    for root, _, files in os.walk(METADATA_PATH):
        for file in files:
            if file.endswith(".json"):
                file_path = os.path.join(root, file)
                try:
                    with open(file_path, "r", encoding="utf-8") as f:
                        json_data = json.load(f)
                        json_list.append(json_data)
                except Exception as e:
                    print(f"Error reading {file_path}: {e}")

    return json_list


try:
    # Note that you can use `client.collections.create_from_dict()` to create a collection from a v3-client-style JSON object
    collection = client.collections.create(
        name="babel",
        vectorizer_config=wvcc.Configure.Vectorizer.text2vec_openai(),
        generative_config=wvcc.Configure.Generative.openai(),
        properties=[
            wvcc.Property(name="category", data_type=wvcc.DataType.TEXT),
            wvcc.Property(name="highlights", data_type=wvcc.DataType.TEXT_ARRAY),
            wvcc.Property(name="keywords", data_type=wvcc.DataType.TEXT_ARRAY),
            wvcc.Property(name="path", data_type=wvcc.DataType.TEXT),
            wvcc.Property(name="references", data_type=wvcc.DataType.TEXT_ARRAY),
            wvcc.Property(name="related_links", data_type=wvcc.DataType.TEXT_ARRAY),
            wvcc.Property(name="summary", data_type=wvcc.DataType.TEXT),
            wvcc.Property(name="tags", data_type=wvcc.DataType.TEXT_ARRAY),
        ],
    )

    babel = client.collections.get("babel")

    # data_rows = get_all_json_files()
    # with babel.batch.dynamic() as batch:
    #     for data_row in data_rows:
    #         batch.add_object(
    #             properties=data_row,
    #         )

    # response = babel.query.fetch_objects()

    # for o in response.objects:
    #     print(o.properties)

    response = babel.generate.near_text(
        query="want all feedbacks",
        limit=2,
        # target_vector="title_country",  # Specify the target vector for named vector collections
        # single_prompt="Translate this into German: {review_body}",
        grouped_task="Summarize these review",
        return_metadata=MetadataQuery(distance=True),
    )

    print(response.generated)
    for o in response.objects:
        print(o.properties)
        print(o.generated)


finally:
    client.close()
