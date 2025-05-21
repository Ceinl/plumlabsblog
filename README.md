# Overview
    Blog engine (Im pretty dump so called it Plublumbs)
    Basicly a web app to parse markdown files to a HTML content and display it as a blog article.
    Use HTMX, Golang, Tailwind CSS and Sqlite. Dont like a lot of overcomplication so yes.
        API and Server make used go http module without any external dependencies.
        Lexer, parser and renderer(I belive that what it is) also written from scratch.

# Features
    
    Now 1 but 2 services, admin and user panel (front with integrated db and back)

    Admin panel
        Backend, standart port 1612 (p and l number orders)
        Have just 4 buttons, to test APIs and thats it

    User Interface 
        FrontEnd, standart port 2113 (u and m number orders)
        List of articles at left and nice goot spot for content in center.

# Usage

Clone project
```
    git clone https://github.com/dmytroslyva/plumlabsblog.git
    cd plumlabsblog
```

Terminal session 1
```
    cd back
    go run main.go
```

Terminal session 2
```
    cd front
    go run main.go
```



