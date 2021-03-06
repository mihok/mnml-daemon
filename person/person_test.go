package person

import "testing"

func TestPerson(t *testing.T) {
    // <setup code>

    t.Run("Person inherits fmt.Stringer", func(t *testing.T) {
      bob := Person{
          "Bob",
          "Bobberson",
      }

      result := bob.String()

      if (result != "Bob Bobberson") {
        t.Error("Expected 'Bob Bobberson', got ", result);
      }
    })

    // <tear-down code>
}
