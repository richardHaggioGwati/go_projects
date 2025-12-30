# RSS AGGREGATOR

The purpose pf the server is to aggregate RSS feeds from various sources. What our server will do is allow our users to add different RSS to it's database, and then it will go and fetch the RSS feed data from those sources periodically. The server will then parse the RSS feed data and store it in a database, so we can view them later on.

## Table of Contents

- [What is RSS Aggregator?](#what-is-rss-aggregator)
  - [Whats is RSS?](#whats-is-rss)
  - [Finally What is RSS Aggregator?](#finally-what-is-rss-aggregator)
- [Tech Stack](#tech-stack)
- [Architecture](#architecture)
- [Getting Started](#getting-started)
- [Features](#features)
- [Installation](#installation)
- [Usage](#usage)
- [Configuration](#configuration)
- [API Endpoints](#api-endpoints)
- [Contributing](#contributing)
- [License](#license)

## What is RSS Aggregator?

### Whats is RSS?

RSS (Really Simple Syndication) is a web feed format used to publish frequently updated content, such as news articles, blog posts, and podcasts. It allows users to subscribe to their favorite websites and receive updates in a standardized format.

An aggregator is a tool that collects and organizes content from multiple RSS feeds into a single location, making it easier for users to access and read the content they are interested in.

### Finally What is RSS Aggregator?

An RSS aggregator is a software application or web service that collects and displays content from multiple RSS feeds in one place. It allows users to subscribe to various RSS feeds and view the latest updates from those sources without having to visit each website individually.

## Tech Stack

- **Backend**: Go
- **Database**: SQL
