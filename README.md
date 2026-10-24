# FeedFolio
FeedFolio is a Golang-based RSS reader that aggregates RSS feeds for users, allowing them to follow their favorite sources and receive updates in a streamlined interface. The backend leverages PostgreSQL, SQLC for type-safe SQL code, and Goose for efficient database migrations.

## Features

- **User Subscriptions**: Users can subscribe to multiple RSS feeds and manage their own list of followed sources.
- **Aggregation Worker**: A background worker periodically fetches and compiles updates from all RSS feeds followed by users, ensuring that new content is available as soon as possible.
- **Type-Safe SQL**: SQLC is used to generate type-safe SQL code, making database queries more maintainable and secure.
- **Database Migrations**: Goose is used to manage and track database migrations, making version control and schema changes more manageable.

## Technology Stack

- **Backend**: Golang
- **Database**: PostgreSQL
- **Migration Tool**: Goose
- **SQL Code Generation**: SQLC

## Usage

- **Adding Feeds**: Users can add RSS feeds to their subscriptions by providing valid RSS feed URLs.
- **Viewing Updates**: Feed items are fetched and stored per user subscription. Users can view the latest updates from their followed feeds.
- **Marking as Read**: Users have the option to mark feed items as read to keep track of new and viewed items.
- **Scheduled Fetching**: FeedFolio includes a scheduled fetcher to pull updates at regular intervals, making new content available without user intervention.
