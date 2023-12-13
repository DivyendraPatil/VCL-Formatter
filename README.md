# VCL Formatter

## Overview

This tool is a VCL Formatter written in Go. It takes a VCL file as input and formats it by adjusting indentation and merging lines where appropriate. The goal is to make VCL code more readable and maintainable.

## Prerequisites

- Go (Golang) installed on your system.

## Usage

To use the VCL Formatter, simply pass the path of your VCL file as an argument:

```
go run main.go <path-to-your-file.vcl>
```

## How It Works

The program reads a VCL file and processes each line:

- Trims whitespace and checks for specific VCL structures like if conditions and sub vcl_ blocks.
- Formats opening braces { by merging them with preceding lines where appropriate.
- Indents the lines based on the nesting level.
- Decreases indentation when closing braces } are encountered.
