## Uitleg opdrachten Applicatie laag 
  
### In dit Bestand zal ik bij elke code een stukje informatie geven over wat deze code doet en waarvoor het gebruikt wordt.  
  
**[Gatekeeper/main.go](gatekeeper/main.go)**  
Er wordt een Config struct gedefinieerd om configuratiegegevens op te slaan.  
Er is ook een Payload struct gedefinieerd om gegevens voor een HTTP-verzoek te bevatten.  
Het programma begint met het openen van logbestanden voor fouten en toegang.  
Het controleert of er een kentekenplaat is opgegeven als argument bij het uitvoeren van het programma. Zo niet, wordt het programma beëindigd.  
Het leest de configuratiegegevens uit een JSON-bestand en decodeert deze in een slice van Config structs.  
Het maakt een HTTP-verzoek voor elke configuratie en voegt het kentekenplaat als een queryparameter toe.  
Het verwerkt de JSON-respons van het verzoek en krijgt de naam van de gebruiker.
Als er geen gebruikersnaam is gevonden, wordt het programma beëindigd.  
Het schrijft toegangsinformatie naar het logbestand en stuurt een HTTP-verzoek naar een API met behulp van de geconfigureerde URL en authenticatiegegevens.  
Het controleert de huidige tijd en geeft berichten weer op basis van de configuratiegegevens.  
Het programma wordt herhaald voor elke configuratie in de slice.  
Kortom, deze code is een programma die toegang verleent tot een parkeerplaats op basis van een kentekenplaat, interactie heeft met een externe API en logboeken bijhoudt.  
    
**[Gatekeeper/gatekeeper.py](gatekeeper/gatekeeper.py)**  
De code importeert de vereiste bibliotheken: cv2, pytesseract en subprocess.  
Er wordt een videostream van de webcam gestart met behulp van de cv2.VideoCapture(0)-functie.  
Er wordt een oneindige lus gestart om continu frames van de videostream te lezen.  
Voor elk frame wordt het volgende gedaan:  
    Het huidige frame wordt gelezen met behulp van cap.read().  
    Het frame wordt omgezet naar grijswaarden met behulp van cv2.cvtColor().  
    Er wordt een Gaussische vervaging toegepast op het grijswaardenbeeld met cv2.GaussianBlur().  
    Randen worden gedetecteerd met behulp van de Canny-randdetectiemethode met cv2.Canny().  
    Contouren worden gevonden in de randafbeelding met behulp van cv2.findContours().  
    De contouren worden gefilterd op basis van grootte om mogelijke kentekenplaten te vinden.  
    Bounding boxen worden getekend rond de gevonden kentekenplaten met behulp van cv2.rectangle().  
    PyTesseract wordt gebruikt voor optische tekenherkenning (OCR) op de individuele kentekenplaatgebieden om de tekens op het kenteken te lezen.  
    Als een kenteken precies 6 tekens bevat, wordt een terminalopdracht uitgevoerd met het kenteken als argument.  
    Het bewerkte frame wordt weergegeven in een venster met cv2.imshow().  
    Als de 'q'-toets wordt ingedrukt, wordt de lus beëindigd.  
Nadat de lus is beëindigd, worden de video-opnamebron vrijgegeven en de vensters gesloten met cap.release() en cv2.destroyAllWindows().  
In het kort, deze code maakt gebruik van computer vision-technieken om kentekenplaten te detecteren en te lezen van een videostream van de webcam, en voert vervolgens een terminalopdracht uit met het gedetecteerde kenteken als argument.  
  
**[config.json](gatekeeper/config.json)**  
"Morning_start_time": De starttijd van de ochtendperiode (7 uur).  
"Noon_start_time": De starttijd van de middagperiode (12 uur).  
"Evening_start_time": De starttijd van de avondperiode (18 uur).  
"No_parking_acces_start_time": De starttijd waarop geen toegang tot de parkeerfaciliteiten is toegestaan (23 uur).  
"API_ip_or_domain": Het IP-adres of domein van de API waarmee wordt gecommuniceerd (http://192.168.137.18).  
"API_port": De poort van de API (80).  
"Morning_message": Het bericht dat wordt weergegeven tijdens de ochtendperiode ("Good morning").  
"Noon_message": Het bericht dat wordt weergegeven tijdens de middagperiode ("Good afternoon").  
"Evening_message": Het bericht dat wordt weergegeven tijdens de avondperiode ("Good evening").  
"No_parking_acces_message": Het bericht dat wordt weergegeven wanneer de parkeerfaciliteiten gesloten zijn ("Sorry, our parking facilities are currently closed.").  
"Technical_dificulties": Het bericht dat wordt weergegeven wanneer er technische problemen zijn ("Sorry, we are currently experiencing technical difficulties").  
"Welcome_message": Het welkomstbericht dat wordt weergegeven ("Welcome at Holiday parks.").  
"Not_allowed": Het bericht dat wordt weergegeven wanneer een kentekenplaat niet is toegestaan ("License plate not permitted!").  
"API_Url": De URL van de API waarmee wordt gecommuniceerd voor het controleren van kentekenplaten (http://4.175.136.232:8080/nummerplaat).  
Deze configuratiegegevens worden gebruikt in de code om berichten weer te geven en communicatie met de API uit te voeren.  
  
**[ESP Gatekeeper](esphome32/gatekeeper.yaml)**  
<a href="/Screenshots/emmer%20vol%20water.PNG"><img src="/Screenshots/emmer%20vol%20water.PNG" align="right" height="90" width="110" ></a>
Deze opdracht vond ik iets uitdagender hieom ook dat deze zig in de Wedstrijd folder bevindt. Ik ben hier persoonlijk wel nog tegen wat dingen aangelopen. Denk hierbij aan hoe zorg ik er voor dat mensen alleen maar getallen in kunnnen vullen, maar ook aan hoe zorg ik dat de annimatie van het vullen soepeltjes verloopt (wat zijn te groten stappen en wat te kleinen). 
  
**[Woordvervormer](Wedstrijd/Woordvervormer/Form1.cs)**  
<a href="/Screenshots/woordvervormer.PNG"><img src="/Screenshots/woordvervormer.PNG" align="right" height="150" width="150" ></a>
Bij deze opdracht liep ik ook weer tegen een aantal puntjes aan, hierom zal ook deze zich bevinden in de wedstrijd folder. De dingen waar ik hier tegen aan liep kwamen voornamelijk voor bij de functies Omdraaien en Oneven. Deze puntjes waar ik tegen aan liep, kwamen vooral doordat ik dingen met een array moest doen die ik eerder nog niet had gedaan. Denk hier vooral aan maak een string met alleen plek 1,3 en 5 van een array, dit had ik nog nooit eerder gedaan en was dus ook echt even uitzoeken. Ook bij het opdraaien van het woord liep ik hier weer even tegenaan, ik heb er uiteindelijk voor gekozen om een nieuwe array te maken. Hierbij 0 beginnend met schrijven en bij het laatste caracter beginnen te lezen bij de oude array, dit zorgt ervoor dat ze precies andersom staan.  
  
**[Blackjack](Wedstrijd/Blackjack/Form1.cs)**  
<a href="/Screenshots/Blackjack.PNG"><img src="/Screenshots/Blackjack.PNG" align="right" height="150" width="300" ></a>
Bij Blackjack heb ik aan het begin echt zitte puzelen, ik wist gewoon echt niet waar/hoe ik moest beginnen. Toen ik eenmaal gevonden had waar ik mee wou beginnen verliep het allemaal vrij soepel, ik had alle UI snel gemaakt en het programma werkte na een paar tweaks vrijwel gelijk zoals verwacht. De rede dat blackjack zich in de Wedstrijd folder bevindt is, dat deze challenge zo is gemaakt dat hij alle hiervoor behandelde onderdelen in een opdracht verwerkt. Hierom was dit ook de meest uitdagende opdracht.
