# finCrawler

<img src="finCrawler.png" width="300">

finCrawler collects data to support more convenient contributions. This project provides scalability by implementing input and output plugins, enabling developers to collect a wider variety of data. The project is developed in the Go language.

## How It Works

----
![finCrawler_explain](https://github.com/finContribution/finCrawler/assets/65060314/2dc662fe-7264-490b-8672-18b458342f81)

### Input

The input module fetches issues and pull requests from remote repository projects. This information is received as a `[]byte` object and stored in memory using the Go channel feature. The output plugin can then retrieve and store this information.

### Output

The output module is responsible for converting the data in the channel into a data type suitable for storage. For example, when using the Elasticsearch module, it converts `[]byte` data into a `_doc` format for indexing. (Uses the Elasticsearch bulk API)

### Others

In addition to the input and output described above, developers can write and use additional modules. Please refer to the upcoming contribute guide for contributions to this project!
