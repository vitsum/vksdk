package api // import "github.com/SevereCloud/vksdk/5.92/api"

import (
	"github.com/SevereCloud/vksdk/5.92/object"
)

// NotesAddResponse struct
type NotesAddResponse int

// NotesAdd creates a new note for the current user.
// https://vk.com/dev/notes.add
func (vk VK) NotesAdd(params map[string]string) (response NotesAddResponse, vkErr Error) {
	vk.requestU("notes.add", params, &response, &vkErr)
	return
}

// NotesCreateCommentResponse struct
type NotesCreateCommentResponse int

// NotesCreateComment adds a new comment on a note.
// https://vk.com/dev/notes.createComment
func (vk VK) NotesCreateComment(params map[string]string) (response NotesCreateCommentResponse, vkErr Error) {
	vk.requestU("notes.createComment", params, &response, &vkErr)
	return
}

// NotesDelete deletes a note of the current user.
// https://vk.com/dev/notes.delete
func (vk VK) NotesDelete(params map[string]string) (vkErr Error) {
	_, vkErr = vk.Request("notes.delete", params)
	return
}

// NotesDeleteComment deletes a comment on a note.
// https://vk.com/dev/notes.deleteComment
func (vk VK) NotesDeleteComment(params map[string]string) (vkErr Error) {
	_, vkErr = vk.Request("notes.deleteComment", params)
	return
}

// NotesEdit edits a note of the current user.
// https://vk.com/dev/notes.edit
func (vk VK) NotesEdit(params map[string]string) (vkErr Error) {
	_, vkErr = vk.Request("notes.edit", params)
	return
}

// NotesEditComment edits a comment on a note.
// https://vk.com/dev/notes.editComment
func (vk VK) NotesEditComment(params map[string]string) (vkErr Error) {
	_, vkErr = vk.Request("notes.editComment", params)
	return
}

// NotesGetResponse struct
type NotesGetResponse struct {
	Count int                `json:"count"`
	Items []object.NotesNote `json:"items"`
}

// NotesGet returns a list of notes created by a user.
// https://vk.com/dev/notes.get
func (vk VK) NotesGet(params map[string]string) (response NotesGetResponse, vkErr Error) {
	vk.requestU("notes.get", params, &response, &vkErr)
	return
}

// NotesGetByIDResponse struct
type NotesGetByIDResponse object.NotesNote

// NotesGetByID returns a note by its ID.
// https://vk.com/dev/notes.getById
func (vk VK) NotesGetByID(params map[string]string) (response NotesGetByIDResponse, vkErr Error) {
	vk.requestU("notes.getById", params, &response, &vkErr)
	return
}

// NotesGetCommentsResponse struct
type NotesGetCommentsResponse struct {
	Count int                       `json:"count"`
	Items []object.NotesNoteComment `json:"items"`
}

// NotesGetComments returns a list of comments on a note.
// https://vk.com/dev/notes.getComments
func (vk VK) NotesGetComments(params map[string]string) (response NotesGetCommentsResponse, vkErr Error) {
	vk.requestU("notes.getComments", params, &response, &vkErr)
	return
}

// NotesRestoreComment restores a deleted comment on a note.
// https://vk.com/dev/notes.restoreComment
func (vk VK) NotesRestoreComment(params map[string]string) (vkErr Error) {
	_, vkErr = vk.Request("notes.restoreComment", params)
	return
}
