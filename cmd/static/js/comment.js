import { session_username } from './auth.js';
import { guid } from './post.js';
import  * as wsoc from './wsManager.js';


// Edit comment textarea hide/show
let edit_comment_buttons = $('.comment-editor-button');
edit_comment_buttons.each(function() {
    $(this).on('click', function() {
        const id = $(this).data('id'); // Use jQuery's .data() method to get the data-id attribute
        const comment_text_id = '#comment-text-' + id;
        const comment_editor_id = '#comment-editor-' + id;
        const edit_form = $(comment_editor_id);
        const comment_text = $(comment_text_id);
        if (edit_form.css('display') === 'none' || edit_form.css('display') === '') {
            edit_form.show();
            comment_text.hide();
        } else {
            edit_form.hide();
            comment_text.show();
        }
    });
});

// Rate comment buttons behaviour
let rate_comment_up_buttons = $('.comment-rate-up-button');
rate_comment_up_buttons.each(function() {
    $(this).on('click', function() {
        const id = $(this).data('id');
        let commentElement = $('.comment-container[data-id="' + id + '"]');
        if (commentElement.length) {
            rateCommentUp(commentElement);
        }
    });
});

let rate_comment_down_buttons = $('.comment-rate-down-button');
rate_comment_down_buttons.each(function() {
    $(this).on('click', function() {
        const id = $(this).data('id');
        let commentElement = $('.comment-container[data-id="' + id + '"]');
        if (commentElement.length) {
            rateCommentDown(commentElement);
        }
    });
});

let rate_comment_hiddens = $('.comment-rating-hidden');
rate_comment_hiddens.each(function() {
    let userRating = $(this).val();
    if (userRating > 0) {
        let rate_up_button = $(this).prev();
        rate_up_button.css('color', 'orange');
    } else if (userRating < 0) {
        let rate_down_button = $(this).next();
        rate_down_button.css('color', 'orange');
    }
});


// Delete comment modal behaviour
const commentModal = $("#modal-container-comment");
let commentIdToDelete = null;
$('.comment-deleter-button').on('click', function() {
    commentIdToDelete = $(this).data('id');
    commentModal.show();
});
$("#delete-comment-sure-no").on('click', function() {
    commentModal.hide();
    commentIdToDelete = null;
});
$("#delete-comment-sure-yes").on('click', function() {
    if (commentIdToDelete !== null) {
        let commentElement = $('.comment-container[data-id="' + commentIdToDelete + '"]');
        if (commentElement.length) {
            deleteComment(commentElement);
        }
        commentModal.hide();
        commentIdToDelete = null;
    }
});
$(window).on('click', function(event) {
    if (event.target === commentModal[0]) {
        commentModal.hide();
        commentIdToDelete = null;
    }
});


// AJAX calls

// Add Comment
$(document).ready(function() {
    $('#add-comment-form').on('submit', function(event) {
        event.preventDefault();
        addComment();
    });
});

function addComment() {
    const formData = $('#add-comment-form').serialize();
    $.ajax({
        url: '/action/post/' + guid + '/comment',
        method: 'POST',
        data: formData,
        success: function(response) {
            location.reload();
        },
        error: function(err) {
            console.error("Error:", err);
        }
    });
}

// Edit Comment
$(document).ready(function() {
    $('.comment-edit-form').on('submit', function(event) {
        event.preventDefault();
        editComment(event.currentTarget);
    });
});

function editComment(formElement) {
    let form = $(formElement);
    let id = form.find('.comment_id').val();
    let edited_comment = form.find('.edit_comment').val();
    $.ajax({
        url: '/action/post/' + guid + '/comment/' + id,
        method: 'PUT',
        data: {
            comment: edited_comment
        },
        success: function() {
            location.reload()
        },
        error: function(err) {
            console.error("Error:", err);
        }
    });
}

// Delete comment
function deleteComment(commentElement) {
    let commentId = commentElement.data('id');
    $.ajax({
        url: '/action/post/' + guid + '/comment/' + commentId,
        method: 'DELETE',
        success: function() {
            location.reload();
        },
        error: function(err) {
            console.error("Error:", err);
        }
    });
}

// Rate comment Up
function rateCommentUp(el) {
    let id = $(el).find('.comment_id').val();
    $.ajax({
        url: '/action/post/' + guid + '/comment/' + id + '/up',
        method: 'POST',
        success: function(res) {
            const message = {
                from_username: session_username,
                type: 'rate_comment',
                msg: wsoc.MSG_COMMENT_RATE,
                resource_id: id,
                parent_id: guid
            };
            wsoc.sendWSmsg(message);
            location.reload()
        },
        error: function(err) {
            console.error("Error:", err);
        }
    })
    return false;
}

// Rate comment Down
function rateCommentDown(el) {
    let id = $(el).find('.comment_id').val();
    $.ajax({
        url: '/action/post/' + guid + '/comment/' + id + '/down',
        method: 'POST',
        success: function(res) {
            const message = {
                from_username: session_username,
                type: 'rate_comment',
                msg: wsoc.MSG_COMMENT_RATE,
                resource_id: id,
                parent_id: guid
            };
            wsoc.sendWSmsg(message);
            location.reload()
        },
        error: function(err) {
            console.error("Error:", err);
        }
    })
    return false;
}