Tender.server("git.imeete.com")(function (config) {
    config.name = 1
    config.b = 2
    config.c = "xxx"
    console.log("这份", this)
})

Tender.server("git.imeete.com")(function (config) {
    config.name = 1
    config.b = 2
    config.c = "xxx"
    console.log("这份", this)
})
