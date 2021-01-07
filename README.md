# FileDiff

A file difference checker written in Go.

## Getting Started

### Prerequisites

You need Go installed on your device in order to install this utility.

### Installing

```
go get github.com/webhooked/filediff
```

## Usage guide

```
filediff file1.css file2.css
```

*Please note that the file checker takes any type of file, meaning it is not limited to JavaScript files as shown in the example below.*

### Example usage

Let's say we have two React components (JS-files) that we want to compare.

#### component1.js
```
// component1.js

function App() {
  return (
    <div className="App">
      <header className="large-header">
        <img src={logo} className="logo" alt="logo" />
        <p>
          Edit <code>src/App.js</code> and save to reload.
        </p>
      </header>
    </div>
  );
}
```

#### component2.js
```
// component2.js

function Home() {
  return (
    <div className="Home">
      <header className="large-header">
        <img src={logo} className="logo" alt="logo" />
        <h1>
          Edit <code>src/Home.js</code> and save to reload.
        </h1>
      </header>
    </div>
  );
}
```

We use the filediff utility to display the differences.

```
filediff component1.js component2.js
```

The utility then provides us with the following result

[Example - React component code differences](url)<img width="478" alt="react" src="https://user-images.githubusercontent.com/9132742/103879762-918c9680-50d8-11eb-87cc-d9eef8f5869e.png">

## License

This project is licensed under the Apache License 2.0 - see the [LICENSE.md](LICENSE) file for details
