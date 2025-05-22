# NS LLMS
[![Go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)](https://go.dev/) [![Python](https://img.shields.io/badge/python-3670A0?style=for-the-badge&logo=python&logoColor=ffdd54)](https://www.python.org/) [![GitHub](https://img.shields.io/badge/GitHub-100000?style=for-the-badge&logo=github&logoColor=white)](https://github.com/)

This is a repository to track the learning process for the 2025 Personal Development Plan in Nicasource. Please note that this repository is not intended to be used as a reference for the course. It is meant to be a tool for tracking the learning process and to provide evidence of the work done. As such, it contains the code for the assignments of the selected courses.

Even though the suggested tech stack for the courses is Python, Go was chosen as the primary language for the assignments. The main motivation for this choice was personal, as I wanted to learn Go with hands-on experience, and I'm already familiar with Python, specially with the tools used in the courses. That said, I'm open to using Python as a fallback if the assignments are too complex for my current knowledge of Go.

## Courses

The following courses were taken during the 2025 PDP:

- [LLM Engineering: Master AI, Large Language Models & Agents](https://www.udemy.com/course/llm-engineering-master-ai-and-large-language-models/)
- [Basic to Advanced: Retreival-Augmented Generation (RAG)](https://www.udemy.com/course/basic-to-advanced-retreival-augmented-generation-rag-course/)
- [Mastering Retrieval-Augmented Generation (RAG)](https://www.udemy.com/course/mastering-retrieval-augmented-generation/)
- [Open-source LLMs: Uncensored & secure AI locally with RAG](https://www.udemy.com/course/open-source-llms-uncensored-secure-ai-locally-with-rag/)

## Code structure

The code for the assignments is organized in the following way:

- /cmd: contains the entrypoints for the assignments (meant to be run from the terminal)
- /shared: contains the code for shared functionality or data structures
- /llm-engineering: contains the code for the LLM Engineering course
- /basic-to-advanced: contains the code for the Basic to Advanced course
- /mastering-rag: contains the code for the Mastering RAG course
- /open-source-llms: contains the code for the Open Source LLMs course

Each course will have its own directory, which will contain the code for the artifacts of the course. The code for the assignments will always be in the cmd directory.

## Running assignments

To run the assignments, you can build or run the code from the entrypoints in /cmd.
