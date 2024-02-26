# catgpt

`NOTE`: If you find this interesting check out [GPTScript](https://github.com/gptscript-ai/gptscript) instead. This is archived as I built that project and I like it better.

Because I'm lazy and this helps me be lazier.

This is a command-line tool that uses the OpenAI model to generate text based on user input.

## Installation

To install `catgpt`, you can run the following command:

```
go install github.com/ibuildthecloud/catgpt@main
```

Or download a binary from the releases.

## Usage

```
echo [MESSAGE] | catgpt [prompt]
```

Options:
- `api-key`: OpenAI API Key (default: environment variable `OPENAI_API_KEY`)
- `base-url`: OpenAI Base URL (default: `https://api.openai.com/v1`)
- `model`: The OpenAI model to use (default: `gpt-3.5-turbo`)

## Example

To use `catgpt` to generate text, you can pipe input from a file or from stdin. Here's an example:

```
echo "What is the meaning of life?" | catgpt
```

This will generate a response from the GPT-3.5 Turbo model based on the prompt "What is the meaning of life?".

A more complex example would be how this README was orignally generated

```
find -name "*.go" -print -exec cat {} \; | catgpt "Create a readme for this github project that has this go file content" > README.md
```

Or how I created the first commit message

```
git diff --staged | catgpt -m gpt-4 "Generate a git commit message"
```


## License

This project is licensed under the [Apache License 2.0](LICENSE).

## Contributing

Contributions are welcome! If you would like to contribute to this project, please open an issue or submit a pull request.
