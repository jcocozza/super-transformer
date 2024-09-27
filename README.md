# 🌟 SUPER TRANSFORMER 🌟

![super-transformer](https://media.giphy.com/media/3ohhwHFlzVhXr6XWNa/giphy.gif)

**Super Transformer** is here to TRANSFORM your data, making life easier, one file at a time. Whether it's Parquet, CSV, or JSON, we've got you covered. This isn't your average transformer, this is a SUPER transformer! 💪✨

---

## 💼 What Does It Do?

🌀 **Super Transformer** takes your raw, structured data (like Parquet files) and transforms it into something usable. Whether you’re working with cloud data, local files, or complex formats, this tool is built to handle them all. It’s all about transforming messy data into a format that can be easily processed or analyzed.

### 🚀 Features

- 🛠 **Parquet Parsing**: Seamless Parquet-to-JSON transformation. No more Parquet headaches!
- 🛠 **Excel Parsing**: Seamless Excel-to-JSON transformation.
- 🎛 **Flexible Transformation**: Handles different data formats & provides custom transformations.
- 💨 **Fast and Efficient**: Optimized for performance with Go's concurrency model.
- 🤖 **CLI Interface**: Use it right from your terminal with a snazzy command line interface.
- 🔄 **Format Extensibility**: Support for adding other formats (future-proof!).

---

## 🏗️ Getting Started

### 📦 Installation

First things first, clone the repo and install dependencies:

```bash
git clone https://github.com/jcocozza/super-transformer.git
cd super-transformer
go build
```

### 🎮 Running the Transformer

After installation, you can run the transformer like this:

```bash
./super-transformer path/to/my-file
```

---

## 🤓 Example Usage

Let’s transform some Parquet data into JSON with just one command:

```bash
super-transformer data/mtcars.parquet
```

BOOM! 💥 Your data is ready in seconds.

---


## 🛠 Configuration

You can configure **Super Transformer** to meet your specific needs.

| Option        | Default        | Description                                         |
| ------------- | -------------- | --------------------------------------------------- |
| `-has-header` | false          | Include this if the CSV has a header (CSV ONLY)     |
| `-split-line` | false          | split by lines (PLAIN TEXT ONLY)                    |
| `-rowlimit`   | -1 (unlimited) | Limit the number of rows transformed (PARQUET ONLY) |

---

## 🤖 Future Features

🎯 **Planned Additions**:

- 🚀 **Support for ORC & Avro**: Broader format support.
- 📊 **Schema Validation**: Automatically validate schemas.
- 🧠 **AI-powered transformations**: Because why not make it fancy?

---

## 🖼️ GIF of Awesomeness

![Transforming Data](https://media.giphy.com/media/xT9IgG50Fb7Mi0prBC/giphy.gif)

---

## 🤝 Contributing

We love collaboration! Submit pull requests, open issues, or make feature suggestions.

---

🌟 **Super Transformer** is here to change the way you work with data! Join us on this journey! 🚀
