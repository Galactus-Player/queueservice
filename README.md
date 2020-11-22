# Queue Service

The service handles the task of managing the queue of videos for rooms. It let's users add, remove and modify the ordering of videos in queue, given a room code such as `1234`. Users are also able to query for the entire queue. The service also does the job of adding metadata such as thumbnail url and title when adding a url of a video to the queue.

The client polls queueservice every few seconds to check if the queue has been updated. The queueservice maintains a simple `counter` for each room. The `counter` is incremented by 1 when a video is added, played (re-ordered queue) or removed. When polling, if the reponse from the service returns a higher counter than the one that's stored on the client, the client is able to update the queue.

## Endpoints

There are 4 simple endpoints. The `room_code` is the number that the roomservice generates.

### Get Queue

```
(GET) /v1/queue/{room_code}
```
Retrieves the list of videos associated with `room_code`. Makes sure to order according to which video is played and the order in which they're added.

### Add Video

```
(POST) /v1/queue/{room_code}/add
```

#### Request Body

```
{
    "url": "<video url>" 
}
```

Adds video to `room_code` and returns the video URL along with the metadata such as video thumbnail and title. Also returns a unique `video_id` for every video so that the client is able to refer to individual videos.

### Remove Video

```
(POST) /v1/queue/{room_code}/remove
```

#### Request Body

```
{
    "id": "<video_id>" 
}
```

Adds video to `room_code` and returns the video URL along with the metadata such as video thumbnail and title. Also returns a unique `video_id` for every video so that the client is able to refer to individual videos.


### Play Video

```
(GET) /v1/queue/{room_code}/play
```

#### Request Body

```
{
    "id": "<video_id>" 
}
```

Changes the priority of video with `video_id` such that it moves to the top of the queue.

## Running the server

To run the server, follow these simple steps:

#### Create Docker network

```
docker network create mongonet
```

#### Run MongoDB

```
docker run -d --network mongonet --name mongo_host -e MONGO_INITDB_ROOT_USERNAME=mongoadmin -e MONGO_INITDB_ROOT_PASSWORD=secret -p 27017:27017 mongo
```

#### Build queueservice image

```
docker build -t queueservice .
```

#### Run queueservice container

```
docker run --network mongonet -p 9090:9090 -it queueservice
```
