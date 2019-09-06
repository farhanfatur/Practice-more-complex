function showData(opt) {
    if (opt == "mongo") {
        $.ajax({
            url: "/api/viewalls",
            type: "POST",
            contentType: "application/json",
            success: function (v) {
                var $tbody = $("tbody");
                $tbody.empty()
                $(".txt-len").hide()
                $.each(v, function(index, val) {
                    var $tr = $("<tr></tr>")
                    $tr.appendTo($tbody)
                    $("<td></td>").html(index+1).appendTo($tr)
                    $("<td></td>").html(val.name).appendTo($tr)
                    $("<td></td>").html(val.age).appendTo($tr)
                    $("<td></td>").html(val.grade).appendTo($tr)
                    $("<td></td>").html("<a href=\"#/edit/"+val._id+"\">Edit</a> | <a href=\"#/delete/"+val._id+"\">Delete</a>").appendTo($tr)
                })
            },
            error: function(i, v) {
            console.log(i, v)
            }
        });
    }else if (opt == "redis") {
        $.ajax({
            url: "/api/keys",
            type: "POST",
            contentType: "application/json",
            success: function (v) {
                var $tbody = $("tbody");
                $tbody.empty()
                $(".txt-len").hide()
                $.each(v, function(index, val) {
                    var $tr = $("<tr></tr>")
                    $tr.appendTo($tbody)
                    $("<td></td>").html(index+1).appendTo($tr)
                    $("<td></td>").html(val).appendTo($tr)
                })
            },
            error: function(i, v) {
            console.log(i, v)
            }
        });
    }
}
function showOneData(id) {
    $("input[type=button]").removeAttr("id")
    $.ajax({
        url: "/api/edits/"+id,
        type: "POST",
        dataType: "json",
        contentType: "application/json",
        success: function(v) {
            $("#name").val(v.name)
            $("#age").val(v.age)
            $("#grade").val(v.grade)
            $(".container-home").append("<input type=\"hidden\" value=\""+v._id+"\" id=\"id\">")
        },
        error: function(v, i) {
            console.log(v, i)
        }
    })
}