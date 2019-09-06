var QueueInsert = Backbone.Model.extend({
    url: "/api/viewalls",
    defaults: {
        number: ""
    },
    parse: function(res) {
        return res
    }
})

var queue = new QueueInsert({})
queue.fetch({
    success: function(data) {
        console.log(JSON.stringify(data))
    }
})