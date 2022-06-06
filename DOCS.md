# API Documentation

Documentation for API endpoints. Also available as a Swagger-based pages, which can be visited at [TODO: Swagger link]().

## **Public Endpoints**

### *Register*

Method: `POST`

`/register`

### *Login & Logout*

Method: `POST`

`/login`

### *Forget Password*

Method: `POST`

- `/reset` for requesting email verficiation code
- `/reset/code` for verifying 6 digit code
- `/reset/newpassword` for applying new password

### *Get All Topics*

Method: `GET`

`/topic`

### *Get Topic's Thread*

Method: `GET`

Available parameters:
- sort_by={best,vote,date}

`/topic/:id`

### *Get Threads*

Method: `GET`

Available parameters:
- sort_by={best,vote,date}
- topic={topic_name/topic_id}

`/thread`

### *Get Thread Details*

Method: `GET` 

Available parameters (for comments):
- sort_by={vote,date}

`/thread/:id`

### *Get Comment Details*

Method: `GET`

Available parameters:
- sort_by={vote,date}

`/thread/:id/comment/:id`

### *Search*

Method: `GET`

Available parameters:
- query={query string}
- topic={topic/topic id}
- sort_by={vote,date}
- limit={number}
- cursor={post id}

### *Check If Username Exist*

Method: `GET`

`/user/check`

## **User Endpoints**

Same as `public`, with additional endpoints:

### *Get Notifications*

Method: `GET`

`/notification`

### *Mark Notification as Read*

Method: `POST`

`/notification/:id`

### *Get/Update Personal Profiles*

Method:
- `GET` - Get personal profile
- `PUT` - Update personal profile

`/profile`

### *Change Password*

Required data:
- Old password (for confirmation)
- New password

Method:
- `PUT`

`/profile/password`

### *Get Other User's Profile*

Method: `GET`

`/user/:id`

### *Create Topic*

Method: `POST`

`/topic/create`

### *Report Thread*

Method: `PUT`

`/thread/:id`

### *Report Comment*

Method: `PUT`

`/thread/:thread_id/comment/:id`

### *Check If Topic Already Exist*

Method: `GET`

`/thread/check`

### *Follow User*

Method: `PUT`

`/user/:id/follow`

### *Unfollow User*

Method: `PUT`

`/user/:id/unfollow`

### *Subscribe Topic*

Method: `PUT`

`/topic/:id/subscribe`

### *Unsubscribe Topic*

Method: `PUT`

`/topic/:id/unsubscribe`

## **Moderator Endpoints**

Same as `user`, with additional endpoints:

### *Update Topic*

Method: `PUT`

`/topic/:id`

### *Approve/Reject Topic-specific Thread/Comment Ban Request*

Method:
- `PUT` for *approve*
- `DELETE` for *reject*

`/moderation/:topic_id/thread/ban_request/:id`

### *Unban Topic-specific Thread/Comment*

Method: `PUT`

`/moderation/:topic_id/thread/banned/:id`

## **Admin Endpoints**

### *Get All User Ban Request*

Method: `GET`

`/moderation/user/ban_request`

### *Get All Banned User*

Method: `GET`

`/moderation/user/banned`

### *Approve/Reject User Ban Request*

Method: 
- `PUT` for *approve*
- `DELETE` for *reject*

`/moderation/user/ban_request/:id`

### *Unban User*

Method: `PUT`

`/moderation/user/banned/:id`

### *Get All Topic Ban Request*

Method: `GET`

`/moderation/topic`

### *Approve/Reject Topic Ban Request*

Method:
- `PUT` for *approve*
- `DELETE` for *reject*

`/moderation/topic/ban_request/:id`

### *Unban Topic*

Method: `PUT`

`/moderation/topic/banned/:id`

### *Approve/Reject Thread/Comment Ban Request*

Method:
- `PUT` for *approve*
- `DELETE` for *reject*

`/moderation/thread/ban_request/:id`

### *Unban Thread/Comment*

Method: `PUT`

`/moderation/thread/banned/:id`
