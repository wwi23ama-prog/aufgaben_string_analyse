package stringhelpers

// ShorterLength erwartet zwei Strings und gibt die Länge des kürzeren zurück.
func ShorterLength(a, b string) int {
	/* Hinweis
	   Benutzen Sie die Funktion len in einer If-Anweisung.
	   Falls a kürzer ist als b, geben Sie die Länge von a zurück,
	   ansonsten die Länge von b.
	*/
	// TODO
	return 0
}

// LongerLength erwartet zwei Strings und gibt die Länge des längeren zurück.
func LongerLength(a, b string) int {
	/* Hinweis
	   Gehen Sie analog zu ShorterLength vor.
	*/
	// TODO
	return 0
}

// DifferenceCount erwartet zwei Strings und gibt die Anzahl der
// Stellen zurück, an denen sich die Strings unterscheiden.
// Falls die Strings unterschiedlich lang sind,
// wird mindestens die Längendifferenz zurückgegeben.
func DifferenceCount(a, b string) int {
	var count int
	/* Hinweis
	   Benutzen Sie die Funktionen ShorterLength und LongerLength.
	   Initialisieren Sie count mit der Längendifferenz
	   und laufen Sie dann in einer Schleife über die kürzere Länge.
	   Vergleichen Sie die Zeichen an der aktuellen Position
	   und falls sie sich unterscheiden, erhöhen Sie count.
	*/
	// TODO
	return count
}

// PositionIsValid erwartet einen String und eine Position.
// Es gibt true zurück, falls die Position innerhalb des Strings liegt.
func PositionIsValid(s string, pos int) bool {
	/* Hinweis
	   Prüfen Sie, ob pos größer oder gleich 0 ist und kleiner als die Länge des Strings ist.
	*/
	// TODO
	return false
}

// SwitchPositions erwartet einen String und zwei Positionen.
// Es gibt einen String zurück, in dem die Zeichen an den beiden
// Positionen vertauscht sind.
// Falls eine der Positionen außerhalb des Strings liegt,
// wird der ursprüngliche String zurückgegeben.
func SwitchPositions(s string, pos1, pos2 int) string {
	/* Hinweis
	   Konvertieren Sie den String in ein Byte-Array.
	   Z.B. mit b := []byte(s)
	   In diesem Array können Sie dann die Zeichen an den Positionen vertauschen
	   und es als String zurückgeben.
	*/
	// TODO
	return ""
}
