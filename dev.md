# install node js
    https://www.digitalocean.com/community/tutorials/how-to-install-node-js-on-ubuntu-20-04-ja
    used apt install
    >> must use option2 otherwise old version

# install postgres by apt
    > made user mayasql in postgres
    > password is not set, check more later

    > deleted mayasql
    # created user , from 
    postgres@maya-jp:~$ createuser maya -P
    maya : coding1000
    # create new database
    CREATE DATABASE go_movies OWNER maya;

# create react app
    npx create=react-app test-app

# install bootstrap
    npm install bootstrap

# OK NOW moving to next APP

# install router
    npm install react-router-dom


# HAD hard issue that can not access the url params in class component
solved by > https://stackoverflow.com/questions/58548767/react-router-dom-useparams-inside-class-component
 
# React router doc
https://reactrouter.com/docs/en/v6/getting-started/tutorial

# react v6 how to pass the params in Link
https://stackoverflow.com/questions/64782949/how-to-pass-params-into-link-using-react-router-v6