# Project Driver
--------------

* Запускать с готовой папкой "loggs"

1. /api/user      -- Get    - Response "200 OK"     - return list of all users
2. /api/user      -- POST   - Response "201 create" - create new user with Username and Password
3. /api/user/{id} -- Ger    - Response "200 OK"     - return one user by ID
4. /api/user/{id} -- Delete - Response "200 OK"     - delete one user by ID   

## TODO:
### API:
1. /api/user/{id} -- Update - Response "204 no content" - update user by ID

### Business logic:
1. Create authorization.
2. Check and validate password.
3. Create table "routes for drivers."
4. Relation table routes and drivers with time start, time finish and photo.
5. Create clients for telegram and google spreadsheets.
