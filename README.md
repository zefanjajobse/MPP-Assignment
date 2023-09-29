De code is deels voorbereid op wat komt, de volgende assessment vraagt om hetzelfde te doen in een rest-api, dus alles van de sub-module movies.go kan grotendeels overgenomen worden als een soort van class.
De huidige error check in de code is om alle errors naar de console sturen en verder niks wordt gedaan, bij een normaal programma zou ik per bekende error een duidelijke error teruggeven en anders de line nummer automatisch laten bijzetten waar het fout gaat. Maar zoals in het mud paper: goed genoeg voor huidige implementatie.
Daarnaast is de code erg basic geschreven door de kleine hoeveelheid ervaring binnen de programmeertaal.

De gekozen restapi is de eerstgevonden oplossing, grote kans dat hier betere oplossingen voor zijn maar zie hier geen reden om hier uitgebreid onderzoek te doen naar de beste oplossing.
Go laat niet toe om meerdere go bestanden in een map te zetten, hierdoor zijn de argumentshander en restapi in een aparte map, zelf vind ik dit minder overzichtelijk dan map per doel.
Voor JSON naar movie convert gebruik ik een techniek gevonden in een andere programmeertaal; rust, blijkbaar kan dat in go ook.
Ook voor de restapi is er geen extra error checking gedaan naast de minimale eis.

Gebaseerd op feedback alsnog meerdere go bestanden in een bestand kunnen zetten, blijkbaar is het een package per map. Maar nu volgt het dezelfde codeerstyle als een van mijn rust projecten.
De summary werkt met een max van 25 movies, want dat zou volgens alle movies binnen codegrade moeten zijn en zodat lokaal getest kan worden zonder verdere aanpassingen/het limiet van api aanvragen bereikt wordt. De workers hebben geen limiet gekregen en draaien zodra het kan, met <-result om te wachten tot alle workers klaar zijn. De struct van de omdb api is van de hele response omdat ik deze keer niet vooruit heb gekeken en hoop dat het gebruikt wordt.
Error checking is hier het minimale aangezien het niet nodig is voor de opdracht.
