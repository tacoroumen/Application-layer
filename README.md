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
Deze code is een configuratiebestand voor het ESPHome-platform, waarmee je ESP32-microcontrollers kunt programmeren en configureren.  

esphome: Dit geeft het begin van de ESPHome-configuratie aan en definieert de naam van het apparaat (gatekeeper).  
logger: Hier wordt de logging geconfigureerd. Met level: VERY_VERBOSE worden zeer gedetailleerde logberichten ingeschakeld.  
esp32: Hier worden de instellingen voor de ESP32-microcontroller gespecificeerd, zoals het gebruikte board (esp32doit-devkit-v1) en het framework (arduino).  
wifi: Dit configureert de wifi-instellingen voor het apparaat, waarbij de ssid (netwerknaam) en het wachtwoord worden gelezen uit geheime waarden (secrets).  
ota: Dit stelt over-the-air (OTA) updates in en specificeert het wachtwoord voor beveiligde updates, ook gelezen uit geheime waarden.  
servo: Hier wordt een servo-motor geconfigureerd met een ID (my_servo) en een uitvoerkanaal (pwm_out).  
output: Dit definieert de uitvoerinstellingen voor het apparaat, in dit geval een ledc-uitvoer (pwm_out) op pin 25 met een frequentie van 50 Hz.  
sensor: Hier wordt een sensor geconfigureerd om de uptime (tijd dat het apparaat actief is) bij te houden.  
web_server: Dit configureert een webserver op poort 80 en stelt authenticatie in met een gebruikersnaam en wachtwoord, gelezen uit geheime waarden. Ook word met de web_server de ESPHome RestAPI aangezet.  
switch: Hier worden schakelaars (actuatoren) geconfigureerd. Er is een GPIO-schakelaar (gate) op pin 26 die wordt gebruikt om een poort te activeren. Er zijn twee acties gedefinieerd voor het in- en uitschakelen van de poort, waarbij ook een servo-motor wordt aangestuurd en een vertraging wordt toegepast. Er is ook een herstartschakelaar (reboot) geconfigureerd.  

Kort samengevat, deze code configureert een ESP32-apparaat met wifi-connectiviteit, een servo-motor, een sensor voor het bijhouden van de uptime en een api met authenticatie. Er zijn ook schakelaars geconfigureerd om een poort te activeren en het apparaat te herstarten.  
  
**[Gatekeeper API](gatekeeper%20API/main.go)**  

  
**[Gatekeeper API/dockerfile](gatekeeper%20API/dockerfile)**  

