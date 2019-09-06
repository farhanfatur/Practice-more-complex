var pushView = Backbone.View.extend({
    render: function() {
        var markup = `<div class="container-home">
        <div class="row">
            <h1>Add Data</h1>
        </div>
        <div class="row">
        <div class="col-sm-6">
            <label>Number</label>
            <input type="number" name="number" id="number" class="form-control">
        </div>
        </div>
        <div class="row">
        <div class="col-sm-6">
            <input type="button" value="submit" id="submitPush" class="btn btn-primary" style="margin-top: 2px;">
        </div>
        </div>
        </div>`
        this.$el.html(markup)
        return this
    },
    events:  {
        "click #submitPush": "doPush"
    },
    doPush: function(event) {
        event.preventDefault()
        var num = {
            "number": $("#number").val()
        }
        var jsonNum = JSON.stringify(num)
        if (num == "") {
            alert("Field is empty")
        }else {
            $.ajax({
                url: "/api/pushs",
                type: "POST",
                data: jsonNum,
                contentType: "application/json",
                success: function(v, i) {
                    view.HomeView.render().trigger("change")
                    $("#number").val("")
                    // showData("redis")
                },
                error: function(v, i) {
                    console.log(v, i)
                }
            })
        }
    }
})