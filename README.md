# adibot: run a bot that talks like you

This is a weekend experiment in which I finetuned a 7B 4-bit quantized LLM on my Discord messages.

The end result was a fairly convincing and pretty hilarious toy that I shared with a few friends.

I published a blog post with what I learned [here](https://adibytes.dev/writing/training-a-bot-on-myself/).

## project structure

- `./discord-bot`: a small discord bot wrapper in Go
- `./model-training`: a set of Python scripts / notebooks for data prep and finetuning the QLoRA adapter

## support

none :)

unfortunately I don't have the bandwidth for feature requests or support questions!
I will, however, accept PRs of any kind, which will be reviewed on a "best-effort" basis.
