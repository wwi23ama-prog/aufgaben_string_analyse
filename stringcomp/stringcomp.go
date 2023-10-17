package stringcomp

import (
	"strings"

	"github.com/wwi23ama-prog/aufgaben_string_analyse/stringhelpers"
)

// AtMostNDifferences erwartet zwei Strings und eine Zahl n.
// Es gibt true zurück, falls die Strings an höchstens n Stellen
// voneinander abweichen.
// Falls die Strings unterschiedlich lang sind,
// gibt es mindestens so viele Abweichungen, wie die Längendifferenz beträgt.
func AtMostNDifferences(a, b string, n int) bool {
	/* Hinweis
	   Benutzen Sie die Funktion DifferenceCount.
	*/
	return stringhelpers.DifferenceCount(a, b) <= n
}

// StartsWith erwartet zwei Strings a, bund gibt true zurück,
// falls b ein Präfix von a ist. Also wenn a mit b beginnt.
func StartsWith(a, b string) bool {
	/* Hinweis
	   Falls b länger als a ist, kann es kein Präfix von a sein.
	   Ansonsten ist b genau dann ein Präfix von a,
	   wenn man durch b laufen kann und dabei das aktuelle Zeichen von b
	   an der selben Stelle auch in a findet.

	   Schreiben Sie also eine Schleife, die über die Positionen in b läuft.
	   Vergleichen Sie in jedem Schritt das Zeichen von b mit dem Zeichen von a
	   an der gleichen Position. Wenn Sie einen Unterschied finden, können
	   Sie vorzeitig false zurückgeben.
	*/
	if len(b) > len(a) {
		return false
	}
	for i := 0; i < len(b); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// EndsWith erwartet zwei Strings a, bund gibt true zurück,
// falls b ein Suffix von a ist. Also wenn a mit b endet.
func EndsWith(a, b string) bool {
	/* Hinweis:
	   Analog zu StartsWith, aber in a darf nicht am Anfang,
	   sondern ab der Position len(a)-len(b) verglichen werden.
	*/
	if len(b) > len(a) {
		return false
	}
	starta := len(a) - len(b)
	for i := 0; i < len(b); i++ {
		if a[starta+i] != b[i] {
			return false
		}
	}
	return true
}

// Contains erwartet zwei Strings a,b und gibt true zurück,
// falls b in a enthalten ist.
func Contains(a, b string) bool {
	/* Hinweis
	   Falls b länger als a ist, kann es nicht in a enthalten sein.
	   Ansonsten schneiden Sie so lange Zeichen vom Anfang von a ab,
	   bis es mit b anfängt oder leer ist.
	*/
	if len(b) > len(a) {
		return false
	}
	for len(a) >= len(b) {
		if StartsWith(a, b) {
			return true
		}
		a = a[1:]
	}
	return false
}

// EqualCaseInsensitive erwartet zwei Strings und gibt true zurück,
// falls sie bis auf Groß- und Kleinschreibung gleich sind.
func EqualCaseInsensitive(a, b string) bool {
	/* Hinweis
	   Sie können die Funktion ToLower aus dem Paket strings benutzen.
	   Beachten Sie auch die Hinweise, die die Umgebung zu Ihrer Lösung gibt.
	   Es gibt ggf. noch eine bessere Lösung als die hier vorgeschlagene.
	*/
	return strings.EqualFold(a, b)
	// Alternative passend zu Hinweis:
	// return strings.ToLower(a) == strings.ToLower(b)
}

// EqualExceptTransposedPositions erwartet zwei Strings a,b und gibt true zurück,
// falls a und b bis auf einen einzelnen Buchstabendreher gleich sind.
func EqualExceptTransposedPositions(a, b string) bool {
	/* Hinweis
	   Falls a und b unterschiedlich lang sind, können sie nicht bis auf
	   einen Buchstabendreher gleich sein.

	   Ansonsten ist a genau dann bis auf einen Buchstabendreher gleich b,
	   wenn es höchstens eine Position gibt, an der man zwei Zeichen
	   vertauschen kann, um a in b zu verwandeln.

	   Schreiben Sie also eine Schleife, die über die Positionen in a läuft.
	   Vertauschen Sie in jedem Schritt die Zeichen an der aktuellen Position
	   und der nächsten Position und vergleichen Sie das Ergebnis mit b.
	*/
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a)-1; i++ {
		switched := stringhelpers.SwitchPositions(a, i, i+1)
		if switched == b {
			return true
		}
	}
	return false
}

// ContainsScatteredSubstring erwartet zwei Strings a,b und gibt true zurück,
// falls der b String ein "verstreuter" Teilstring von a ist.
// Das heißt, dass die Zeichen in b in der gleichen Reihenfolge
// iwie in a vorkommen, aber nicht unbedingt direkt hintereinander.
// Anders ausgedrückt: b geht aus a hervor, indem man einige (oder alle)
// Zeichen streicht.
func ContainsScatteredSubstring(a, b string) bool {
	/* Hinweis
	   Falls b länger als a ist, kann es nicht verstreut in a enthalten sein.
	   Dies sollte als erstes geprüft werden.

	   Ansonsten ist b genau dann ein verstreuter Teilstring von a,
	   wenn man durch b laufen kann und dabei alle Zeichen von a in der gleichen
	   Reihenfolge findet.

	   Man läuft also durch beide Strings gleichzeitig,
	   z.B. mit einer Schleife, die über die Positionen in b läuft.
	   Zusätzlich führt man auch die Position in a mit.
	   Nach jedem Schritt in b macht man so viele Schritte in a,
	   bis man dort das gleiche Zeichen wie in b gefunden hat.
	   Wenn man auf diese Weise tatächlich bis zum Ende von b kommt,
	   ist b ein verstreuter Teilstring von a.
	   Wenn man aber vorzeitig das Ende von a erreicht,
	   ist b kein verstreuter Teilstring von a.
	*/
	if len(b) > len(a) {
		return false
	}
	posa := 0
	posb := 0
	for posb < len(b) {
		// Falls a zu Ende ist, ist b kein verstreuter Teilstring von a.
		if posa >= len(a) {
			return false
		}
		// Finde nächstes Zeichen von b in a.
		for a[posa] != b[posb] {
			posa++
			// Falls a zu Ende ist, ist b kein verstreuter Teilstring von a.
			if posa == len(a) {
				return false
			}
		}
		posa++
		posb++
	}

	return true
}
