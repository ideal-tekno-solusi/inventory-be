# inventory-be

TODO:
- [x] create base code template
- [x] create working database connection
- [x] create example 1 endpoint
- [x] create working endpoint to show all item in inventory
- [x] create working endpoint to show category
- [x] create working endpoint to insert category
- [x] create working endpoint to update category
- [x] create working endpoint to delete category
- [x] create generate code verifier for oauth 2.0 flow https://datatracker.ietf.org/doc/html/rfc7636#section-4.1
- [x] create generate code challenge for oauth 2.0 flow https://datatracker.ietf.org/doc/html/rfc7636#section-4.2
- [x] update flow login for oauth pkce https://youtu.be/nyjEDSGwN1o?si=i2QxWTFoxYzjbZT9
- [x] create endpoint callback to req token
- [x] update validator to using validate playground
- [ ] create middleware to validate auth token
- [ ] create working endpoint to login (redirect to sso)
- [ ] create working endpoint to show location
- [ ] create working endpoint to insert location
- [ ] create working endpoint to update location
- [ ] create working endpoint to delete location
- [ ] create working endpoint to show position
- [ ] create working endpoint to insert position
- [ ] create working endpoint to update position
- [ ] create working endpoint to delete position
- [ ] create working endpoint to show global item
- [ ] create working endpoint to insert global item
- [ ] create working endpoint to update global item
- [ ] create working endpoint to delete global item
- [ ] create working endpoint to insert new item
- [ ] implement swagger

# note
- to run this project in debug with profile staging, use `mvn spring-boot:run -Dspring-boot.run.profiles=staging`
- params for using other profile is different whe running on terminal and in vscode, please be aware