Tender.server("git.imeete.com")(function (server) {
    server.name = 1
    server.b = 2
    server.c = "xxx"
    console.log("这份", this)
})

Tender.server("git.imeete.com")(function (server) {
    server.name = 1
    server.b = 2
    server.c = "xxx"
    console.log("这份", this)

    server.tls()(function (tls) {

    })
})