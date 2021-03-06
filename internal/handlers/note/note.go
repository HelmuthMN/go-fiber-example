package noteHandler

import (
	"github.com/HelmuthMN/go-fiber-example/database"
	"github.com/HelmuthMN/go-fiber-example/internal/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type Note struct {
	ID       uuid.UUID `json:"id"`
	Title    string    `json:"title"`
	SubTitle string    `json:"sub_title"`
	Text     string    `json:"text"`
}

func createNoteResponse(noteModel models.Note) Note {
	return Note{ID: noteModel.ID, Title: noteModel.Title, SubTitle: noteModel.SubTitle, Text: noteModel.Text}
}

func GetNotes(c *fiber.Ctx) error {
	db := database.DB
	notes := []models.Note{}

	// find all notes in the database
	db.Find(&notes)

	// if no note is present return an error

	if len(notes) == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No notes present", "data": nil})
	}

	responseNotes := []Note{}
	for _, note := range notes {
		responseNote := createNoteResponse(note)
		responseNotes = append(responseNotes, responseNote)
	}
	return c.Status(200).JSON(responseNotes)
}

func CreateNotes(c *fiber.Ctx) error {
	db := database.DB
	var note models.Note

	// Store the body in the note and return error if encountered
	err := c.BodyParser(&note)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	// Add a uuid to the note
	note.ID = uuid.New()
	// Create the Note and return error if encountered
	err = db.Create(&note).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create note", "data": err})
	}

	responseNote := createNoteResponse(note)
	return c.Status(200).JSON(responseNote)
	// Return the created note
	// return c.JSON(fiber.Map{"status": "success", "message": "Note Created", "data": note})
}

func GetNote(c *fiber.Ctx) error {
	db := database.DB
	var note models.Note

	// Read the param noteID
	id := c.Params("noteId")

	// Find the note with the given ID
	db.Find(&note, "id = ?", id)

	if note.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No note present", "data": nil})
	}
	responseNote := createNoteResponse(note)

	return c.Status(200).JSON(responseNote)
	// return c.JSON(fiber.Map{"status": "success", "message": "Note Found", "data": note})
}

func UpdateNote(c *fiber.Ctx) error {
	type updateNote struct {
		Title    string `json:"title"`
		SubTitle string `json:"sub_title"`
		Text     string `json:"text"`
	}

	db := database.DB
	var note models.Note

	// Read the param noteID
	id := c.Params("noteId")

	// Find the note with the given ID
	db.Find(&note, "id = ?", id)

	if note.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No note present", "data": nil})
	}

	// Store the body containing the updated data and return error if encountered
	var updateNoteData updateNote
	err := c.BodyParser(&updateNoteData)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	// Edit Note
	note.Title = updateNoteData.Title
	note.SubTitle = updateNoteData.SubTitle
	note.Text = updateNoteData.Text

	// Save changes
	db.Save(&note)
	responseNote := createNoteResponse(note)

	return c.Status(200).JSON(responseNote)
	// Return the updated note
	// return c.JSON(fiber.Map{"status": "success", "message": "Notes Found", "data": note})
}

func DeleteNote(c *fiber.Ctx) error {
	db := database.DB
	var note models.Note

	// Read the param noteId
	id := c.Params("noteId")

	// Find the note with the given ID
	db.Find(&note, "id = ?", id)

	// If no such note present return an error
	if note.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No note present", "data": nil})
	}

	// Delete the note and return error if encountered
	err := db.Delete(&note, "id = ?", id).Error

	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Failed to delete note", "data": nil})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Deleted Note"})
}
