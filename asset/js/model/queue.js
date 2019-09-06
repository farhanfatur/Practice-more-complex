var keyModel = Backbone.Model.extend({
    url: "/api/keys",
})

var keyModels = new keyModel()
keyModels.fetch({
    success: function (user) {
        keyModels.set(user.attributes)
    }
})

// var QueueCollection = Backbone.Collection.extend({
//     url: "/api/keys",
//     model: keyModel
// })

// var queueCollection = new QueueCollection()
// queueCollection.fetch({
//     success: function(v, i) {
//         console.log(i)
//     },
//     error: function(v, i) {
//         console.log(v, i)
//     }
// })