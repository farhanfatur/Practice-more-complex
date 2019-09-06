var homeView = Backbone.View.extend({
    render: function() {
        var markup = '<div class="container-home table-custom">'+
        '<p class="txt-len">Total : <span id="totalLen"></span></p>'+
        // Queue & Stack(Redis)
        '<button class="btn btn-success" style="margin-bottom: 2px"  id="push">Push</button>'+
        '<button class="btn btn-primary" style="margin-bottom: 2px;margin-left: 2px;" id="len">Len</button>'+
        '<button class="btn btn-warning" style="margin-bottom: 2px;margin-left: 2px;" id="pop">pop</button>'+
        // Crud(Mongo)
        // '<button class="btn btn-success" style="margin-bottom: 2px" onclick=\"window.location.href=\'#/add\' \">Add</button>'+
        '<div class="content-push">'+
        '<table class="table">'+
        '<thead class="thead-dark">'+
        '<tr>'+
        '<th>No</th><th>Number</th>'+
        '</tr>'+
        '</thead>'+
        '<tbody class="tbody-light">'+
        '</tbody>'+
        '</table>'+
        '</div>'+
        '</div>'

        
        this.$el.html(markup)
        showData("redis")
        return this
    },
    events: {
        "click #len": "showLen",
        "click #pop": "popData",
        "click #push": "showPush"
    },
    showPush: function() {
        $(".container-home").removeClass("table-custom")
        $("#push-content").html(view.PushView.render().el)
    },
    showLen: function() {
        $.ajax({
            url: "api/lens",
            type: "POST",
            contentType: "application/json",
            dataType: "json",
            success: function(v, i) {
                $(".txt-len").show()
                $("#totalLen").text(v)
            }
        })
    },
    popData: function() {
        $.ajax({
            url: "/api/pops",
            type: "POST",
            dataType: "json",
            contentType: "application/json",
            success: function(v, i) {
                console.log("Data ", v, " is pop")
                app.navigate("", {trigger: true})
                showData("redis")
            }
        })
    }
})