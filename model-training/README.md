# Model training tools

This folder contains raw tools used to manipulate Discord export data into training/validation sets, as well as finetune a model using the QLoRA technique.

Note that the finetuning technique used here requires an **Apple Silicon Mac** since it uses MLX.

## Data prep

### Step 1: Get Discord message data

[Request a copy of your Discord data.](https://support.discord.com/hc/en-us/articles/360004027692-Requesting-a-Copy-of-your-Data) This will take about a day, and will contain a _lot_ of information.

Once available and downloaded/unzipped, out of your Discord export package, grab the `messages` folder and copy it to this `model-training` folder with the name `raw_data`.

### Step 2: Convert the data into training/validation sets

While in this current `model-training` folder...

```shell
# create data directory
mkdir -p ./data
# install dependencies
poetry install
# execute notebook
jupyter execute
```

You can now inspect the `.jsonl` datasets in `./data`.

> [!TIP]
> You can adjust the message clustering technique by configuring the line in the notebook labeled `pl.duration(minutes=15)`.

### Step 3: Quantize the model

```shell
# open a shell with the venv `source`-d
poetry shell
# convert model
mlx_lm.convert \
  --hf-path mistralai/Mistral-7B-v0.3 \
  -q
```

### Step 4: Train a QLoRA adapter

Pick your favorite number for a seed.

```shell
mlx_lm.lora --train \
  --model ./mlx_model \
  --data ./data \
  --iters 5000 \
  --learning-rate 1e-5 \
  --seed 123456
```

This took about 5 hours on a M1 Pro with 16GB of total memory.

### Step 5: Run a server

MLX ships with a simple OpenAI API-compatible server.

```shell
mlx_lm.server \
  --model ./mlx_model \
  --adapter-path ./adapters
```

To call the server:

```shell
curl -s localhost:8080/v1/completions \
        -H "Content-Type: application/json" \
        -d '{ "prompt": "hello there\n", "max_tokens": 100, "temperature": 0.5 }'
```
