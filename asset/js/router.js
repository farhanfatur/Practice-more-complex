var applicationView = Backbone.View.extend({
    initialize: function() {
        this.PushView = new pushView()
        this.HomeView = new homeView()
    }
})

var appRouter = Backbone.Router.extend({
    routes: {
        "": "index",
        "view": "View"
        // "add": "viewAdd",
        // "push": "viewPush",
        // "edit/:id": "viewEdit"

    },
    index: function() {
        var dataJson = keyModels.attributes
        console.log(JSON.parse(dataJson))
        $("#content").html(view.HomeView.render(dataJson).el)
    }
    // viewPush: function() {
    //     $("#content").html(this.PushView.render().el)
    // },
    // viewAdd: function() {
    //     $("#content").html(this.AddView.render().el)
    //     $("#title").text("Add")
    // },
    // viewEdit: function(id) {
    //    $("#content").html(this.EditView.render(id).el)
    //    $("#title").text("Edit")
    // }
})

var app
var view
$(function() {
    app = new appRouter()
    view = new applicationView()
    Backbone.history.start()
})