import 'package:flutter/material.dart';
import 'dart:convert';
import 'package:dio/dio.dart';

// The backend's URL
const ApiURL = 'https://example.com:8080/';

// The proceeding screen for mobile app
class ProceedingScreen extends StatefulWidget {
  @override
  _ProceedingScreenState createState() => _ProceedingScreenState();
}

class _ProceedingScreenState extends State<ProceedingScreen> {
  // One controller for each candidate
  final idController = TextEditingController();
  final numeroController = TextEditingController();
  final circunscripcionController = TextEditingController();
  final departamentoController = TextEditingController();
  final provinciaController = TextEditingController();
  final municipioController = TextEditingController();
  final localidadController = TextEditingController();
  final recintoController = TextEditingController();
  final ccFirstController = TextEditingController();
  final ccSecondController = TextEditingController();
  final fpvFirstController = TextEditingController();
  final fpvSecondController = TextEditingController();
  final mtsFirstController = TextEditingController();
  final mtsSecondController = TextEditingController();
  final ucsFirstController = TextEditingController();
  final ucsSecondController = TextEditingController();
  final masFirstController = TextEditingController();
  final masSecondController = TextEditingController();
  final vefFirstController = TextEditingController();
  final vefSecondController = TextEditingController();
  final pdcFirstController = TextEditingController();
  final pdcSecondController = TextEditingController();
  final mnrFirstController = TextEditingController();
  final mnrSecondController = TextEditingController();
  final panFirstController = TextEditingController();
  final panSecondController = TextEditingController();
  final validoFirstController = TextEditingController();
  final validoSecondController = TextEditingController();
  final blancoFirstController = TextEditingController();
  final blancoSecondController = TextEditingController();
  final nuloFirstController = TextEditingController();
  final nuloSecondController = TextEditingController();
  final totalController = TextEditingController();

  @override
  void dispose() {
    idController.dispose();
    numeroController.dispose();
    circunscripcionController.dispose();
    departamentoController.dispose();
    provinciaController.dispose();
    municipioController.dispose();
    localidadController.dispose();
    recintoController.dispose();
    ccFirstController.dispose();
    ccSecondController.dispose();
    fpvFirstController.dispose();
    fpvSecondController.dispose();
    mtsFirstController.dispose();
    mtsSecondController.dispose();
    ucsFirstController.dispose();
    ucsSecondController.dispose();
    masFirstController.dispose();
    masSecondController.dispose();
    vefFirstController.dispose();
    vefSecondController.dispose();
    pdcFirstController.dispose();
    pdcSecondController.dispose();
    mnrFirstController.dispose();
    mnrSecondController.dispose();
    panFirstController.dispose();
    panSecondController.dispose();
    validoFirstController.dispose();
    validoSecondController.dispose();
    blancoFirstController.dispose();
    blancoSecondController.dispose();
    nuloFirstController.dispose();
    nuloSecondController.dispose();
    totalController.dispose();
    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    // Extract the arguments
    final ResponseArguments args = ModalRoute.of(context).settings.arguments;
    String response = args.response;
    Map proceedingMap = jsonDecode(response);
    var proceeding = Proceeding.fromJson(proceedingMap);

    // Show data provided by the backend
    idController.text = proceeding.id.toString();
    numeroController.text = proceeding.numero.toString();
    circunscripcionController.text = proceeding.circunscripcion.toString();
    departamentoController.text = proceeding.departamento;
    provinciaController.text = proceeding.provincia;
    municipioController.text = proceeding.municipio;
    localidadController.text = proceeding.localidad;
    recintoController.text = proceeding.recinto;
    ccFirstController.text = proceeding.presidenteCC.toString();
    ccSecondController.text = proceeding.diputadoCC.toString();
    fpvFirstController.text = proceeding.presidenteFPV.toString();
    fpvSecondController.text = proceeding.diputadoFPV.toString();
    mtsFirstController.text = proceeding.presidenteMTS.toString();
    mtsSecondController.text = proceeding.diputadoMTS.toString();
    ucsFirstController.text = proceeding.presidenteUCS.toString();
    ucsSecondController.text = proceeding.diputadoUCS.toString();
    masFirstController.text = proceeding.presidenteMAS.toString();
    masSecondController.text = proceeding.diputadoMAS.toString();
    vefFirstController.text = proceeding.presidenteVeF.toString();
    vefSecondController.text = proceeding.diputadoVeF.toString();
    pdcFirstController.text = proceeding.presidentePDC.toString();
    pdcSecondController.text = proceeding.diputadoPDC.toString();
    mnrFirstController.text = proceeding.presidenteMNR.toString();
    mnrSecondController.text = proceeding.diputadoMNR.toString();
    panFirstController.text = proceeding.presidentePAN.toString();
    panSecondController.text = proceeding.diputadoPAN.toString();
    validoFirstController.text = proceeding.presidenteValido.toString();
    validoSecondController.text = proceeding.diputadoValido.toString();
    blancoFirstController.text = proceeding.presidenteBlanco.toString();
    blancoSecondController.text = proceeding.diputadoBlanco.toString();
    nuloFirstController.text = proceeding.presidenteNulo.toString();
    nuloSecondController.text = proceeding.diputadoNulo.toString();
    totalController.text = proceeding.votosTotal.toString();

    final AlertDialog loadingDialog = AlertDialog(
      title: Text('Procesando datos'),
      content: Column(
        mainAxisAlignment: MainAxisAlignment.center,
        mainAxisSize: MainAxisSize.min,
        children: [
          SizedBox(
            width: 24,
            height: 24,
            child: CircularProgressIndicator(
              valueColor: AlwaysStoppedAnimation(Color(0xFF0e3e3f)),
            ),
          )
        ],
      ),
    );

    final AlertDialog okDialog = AlertDialog(
      title: Text('Enviar datos'),
      content: Text('Datos enviados correctamente.'),
      actions: [
        FlatButton(
          textColor: Color(0xFF0e3e3f),
          onPressed: () {
            Navigator.of(context).pushReplacementNamed('/home');
          },
          child: Text('ACEPTAR'),
        )
      ],
    );
    final AlertDialog updateDialog = AlertDialog(
      title: Text('Actualizar datos'),
      content: Text(
          'Esta acta ya se encuentra en el servidor. ¿Desea actualizar estos datos?'),
      actions: [
        FlatButton(
          textColor: Color(0xFF0e3e3f),
          onPressed: () {
            Navigator.pop(context);
          },
          child: Text('CANCELAR'),
        ),
        FlatButton(
          textColor: Color(0xFF0e3e3f),
          onPressed: () async {
            Navigator.pop(context);
            Proceeding data = Proceeding(
                int.parse(idController.text),
                int.parse(numeroController.text),
                int.parse(circunscripcionController.text),
                departamentoController.text,
                provinciaController.text,
                municipioController.text,
                localidadController.text,
                recintoController.text,
                int.parse(ccFirstController.text),
                int.parse(fpvFirstController.text),
                int.parse(mtsFirstController.text),
                int.parse(ucsFirstController.text),
                int.parse(masFirstController.text),
                int.parse(vefFirstController.text),
                int.parse(pdcFirstController.text),
                int.parse(mnrFirstController.text),
                int.parse(panFirstController.text),
                int.parse(validoFirstController.text),
                int.parse(blancoFirstController.text),
                int.parse(nuloFirstController.text),
                int.parse(ccSecondController.text),
                int.parse(fpvSecondController.text),
                int.parse(mtsSecondController.text),
                int.parse(ucsSecondController.text),
                int.parse(masSecondController.text),
                int.parse(vefSecondController.text),
                int.parse(pdcSecondController.text),
                int.parse(mnrSecondController.text),
                int.parse(panSecondController.text),
                int.parse(validoSecondController.text),
                int.parse(blancoSecondController.text),
                int.parse(nuloSecondController.text),
                proceeding.dispositivo,
                proceeding.ubicacion,
                int.parse(totalController.text));
            String json = jsonEncode(data);
            // Update data in backend
            try {
              Dio dio = new Dio();
              await dio.post(ApiURL + "api/v1/sw1/send-update-data",
                  data: json);
              showDialog(context: context, builder: (context) => okDialog);
            } on DioError catch (e) {
              print(e);
            }
          },
          child: Text('ACEPTAR'),
        )
      ],
    );

    return Scaffold(
      appBar: AppBar(
        // Here we take the value from the MyHomePage object that was created by
        // the App.build method, and use it to set our appbar title.
        title: Text("VERIFICAR DATOS"),
      ),
      body: Row(
        children: [
          Expanded(
            child: ListView(
              children: [
                Row(
                  children: [
                    Expanded(
                        child: Padding(
                            padding: EdgeInsets.fromLTRB(16.0, 16.0, 16.0, 0.0),
                            child: Text(
                                'Para realizar un mejor conteo, antes de enviar los datos al servidor por favor verifica que los datos extraídos de la imagen sean correctos.',
                                textAlign: TextAlign.center,
                                style: TextStyle(color: Color(0xFF0e3e3f))))),
                  ],
                ),
                Row(
                  children: [
                    Expanded(
                        child: Padding(
                            padding: EdgeInsets.fromLTRB(16.0, 16.0, 16.0, 0.0),
                            child: Text('Código de mesa:',
                                style: TextStyle(
                                    fontSize: 14,
                                    fontWeight: FontWeight.w800)))),
                  ],
                ),
                Row(
                  children: [
                    Expanded(
                        child: Padding(
                      padding: EdgeInsets.fromLTRB(16.0, 8.0, 16.0, 8.0),
                      child: TextFormField(
                        controller: idController,
                        decoration: InputDecoration(
                          labelText: 'Código de mesa',
                        ),
                      ),
                    )),
                    Expanded(
                        child: Padding(
                      padding: EdgeInsets.fromLTRB(16.0, 8.0, 16.0, 8.0),
                      child: TextFormField(
                        controller: numeroController,
                        decoration: InputDecoration(
                          labelText: 'Número de mesa',
                        ),
                      ),
                    ))
                  ],
                ),
                Row(
                  children: [
                    Expanded(
                        child: Padding(
                      padding: EdgeInsets.fromLTRB(16.0, 8.0, 16.0, 8.0),
                      child: TextFormField(
                        controller: circunscripcionController,
                        decoration: InputDecoration(
                          labelText: 'Cir. Uninominal',
                        ),
                      ),
                    )),
                    Expanded(
                        child: Padding(
                            padding: EdgeInsets.fromLTRB(16.0, 16.0, 16.0, 0.0),
                            child: Text(
                              '',
                            ))),
                  ],
                ),
                Row(
                  children: [
                    Expanded(
                        child: Padding(
                            padding: EdgeInsets.fromLTRB(16.0, 16.0, 16.0, 0.0),
                            child: Text('Ubicación de la mesa:',
                                style: TextStyle(
                                    fontSize: 14,
                                    fontWeight: FontWeight.w800)))),
                  ],
                ),
                Row(
                  children: [
                    Expanded(
                        child: Padding(
                      padding: EdgeInsets.fromLTRB(16.0, 8.0, 16.0, 8.0),
                      child: TextFormField(
                        controller: departamentoController,
                        decoration: InputDecoration(
                          labelText: 'Departamento',
                        ),
                      ),
                    )),
                  ],
                ),
                Row(
                  children: [
                    Expanded(
                        child: Padding(
                      padding: EdgeInsets.fromLTRB(16.0, 8.0, 16.0, 8.0),
                      child: TextFormField(
                        controller: provinciaController,
                        decoration: InputDecoration(
                          labelText: 'Provincia',
                        ),
                      ),
                    )),
                  ],
                ),
                Row(
                  children: [
                    Expanded(
                        child: Padding(
                      padding: EdgeInsets.fromLTRB(16.0, 8.0, 16.0, 8.0),
                      child: TextFormField(
                        controller: municipioController,
                        decoration: InputDecoration(
                          labelText: 'Municipio',
                        ),
                      ),
                    )),
                  ],
                ),
                Row(
                  children: [
                    Expanded(
                        child: Padding(
                      padding: EdgeInsets.fromLTRB(16.0, 8.0, 16.0, 8.0),
                      child: TextFormField(
                        controller: localidadController,
                        decoration: InputDecoration(
                          labelText: 'Localidad',
                        ),
                      ),
                    )),
                  ],
                ),
                Row(
                  children: [
                    Expanded(
                        child: Padding(
                      padding: EdgeInsets.fromLTRB(16.0, 8.0, 16.0, 8.0),
                      child: TextFormField(
                        controller: recintoController,
                        decoration: InputDecoration(
                          labelText: 'Recinto',
                        ),
                      ),
                    )),
                  ],
                ),
                Row(
                  children: [
                    Expanded(
                        child: Padding(
                            padding: EdgeInsets.fromLTRB(16.0, 16.0, 16.0, 0.0),
                            child: Text('Presidente/a:',
                                style: TextStyle(
                                    fontSize: 14,
                                    fontWeight: FontWeight.w800)))),
                    Expanded(
                        child: Padding(
                            padding: EdgeInsets.fromLTRB(16.0, 16.0, 16.0, 0.0),
                            child: Text('Diputado/a:',
                                style: TextStyle(
                                    fontSize: 14,
                                    fontWeight: FontWeight.w800))))
                  ],
                ),
                Row(
                  children: [
                    Expanded(
                        child: Padding(
                      padding: EdgeInsets.fromLTRB(16.0, 8.0, 16.0, 8.0),
                      child: TextFormField(
                        controller: ccFirstController,
                        decoration: InputDecoration(
                          labelText: 'C.C.',
                        ),
                      ),
                    )),
                    Expanded(
                        child: Padding(
                      padding: EdgeInsets.fromLTRB(16.0, 8.0, 16.0, 8.0),
                      child: TextFormField(
                        controller: ccSecondController,
                        decoration: InputDecoration(
                          labelText: 'C.C.',
                        ),
                      ),
                    ))
                  ],
                ),
                Row(
                  children: [
                    Expanded(
                        child: Padding(
                      padding: EdgeInsets.fromLTRB(16.0, 8.0, 16.0, 8.0),
                      child: TextFormField(
                        controller: fpvFirstController,
                        decoration: InputDecoration(
                          labelText: 'FPV',
                        ),
                      ),
                    )),
                    Expanded(
                        child: Padding(
                      padding: EdgeInsets.fromLTRB(16.0, 8.0, 16.0, 8.0),
                      child: TextFormField(
                        controller: fpvSecondController,
                        decoration: InputDecoration(
                          labelText: 'FPV',
                        ),
                      ),
                    ))
                  ],
                ),
                Row(
                  children: [
                    Expanded(
                        child: Padding(
                      padding: EdgeInsets.fromLTRB(16.0, 8.0, 16.0, 8.0),
                      child: TextFormField(
                        controller: mtsFirstController,
                        decoration: InputDecoration(
                          labelText: 'MTS',
                        ),
                      ),
                    )),
                    Expanded(
                        child: Padding(
                      padding: EdgeInsets.fromLTRB(16.0, 8.0, 16.0, 8.0),
                      child: TextFormField(
                        controller: mtsSecondController,
                        decoration: InputDecoration(
                          labelText: 'MTS',
                        ),
                      ),
                    ))
                  ],
                ),
                Row(
                  children: [
                    Expanded(
                        child: Padding(
                      padding: EdgeInsets.fromLTRB(16.0, 8.0, 16.0, 8.0),
                      child: TextFormField(
                        controller: ucsFirstController,
                        decoration: InputDecoration(
                          labelText: 'UCS',
                        ),
                      ),
                    )),
                    Expanded(
                        child: Padding(
                      padding: EdgeInsets.fromLTRB(16.0, 8.0, 16.0, 8.0),
                      child: TextFormField(
                        controller: ucsSecondController,
                        decoration: InputDecoration(
                          labelText: 'UCS',
                        ),
                      ),
                    ))
                  ],
                ),
                Row(
                  children: [
                    Expanded(
                        child: Padding(
                      padding: EdgeInsets.fromLTRB(16.0, 8.0, 16.0, 8.0),
                      child: TextFormField(
                        controller: masFirstController,
                        decoration: InputDecoration(
                          labelText: 'MAS-IPSP',
                        ),
                      ),
                    )),
                    Expanded(
                        child: Padding(
                      padding: EdgeInsets.fromLTRB(16.0, 8.0, 16.0, 8.0),
                      child: TextFormField(
                        controller: masSecondController,
                        decoration: InputDecoration(
                          labelText: 'MAS-IPSP',
                        ),
                      ),
                    ))
                  ],
                ),
                Row(
                  children: [
                    Expanded(
                        child: Padding(
                      padding: EdgeInsets.fromLTRB(16.0, 8.0, 16.0, 8.0),
                      child: TextFormField(
                        controller: vefFirstController,
                        decoration: InputDecoration(
                          labelText: '21F',
                        ),
                      ),
                    )),
                    Expanded(
                        child: Padding(
                      padding: EdgeInsets.fromLTRB(16.0, 8.0, 16.0, 8.0),
                      child: TextFormField(
                        controller: vefSecondController,
                        decoration: InputDecoration(
                          labelText: '21F',
                        ),
                      ),
                    ))
                  ],
                ),
                Row(
                  children: [
                    Expanded(
                        child: Padding(
                      padding: EdgeInsets.fromLTRB(16.0, 8.0, 16.0, 8.0),
                      child: TextFormField(
                        controller: pdcFirstController,
                        decoration: InputDecoration(
                          labelText: 'PDC',
                        ),
                      ),
                    )),
                    Expanded(
                        child: Padding(
                      padding: EdgeInsets.fromLTRB(16.0, 8.0, 16.0, 8.0),
                      child: TextFormField(
                        controller: pdcSecondController,
                        decoration: InputDecoration(
                          labelText: 'PDC',
                        ),
                      ),
                    ))
                  ],
                ),
                Row(
                  children: [
                    Expanded(
                        child: Padding(
                      padding: EdgeInsets.fromLTRB(16.0, 8.0, 16.0, 8.0),
                      child: TextFormField(
                        controller: mnrFirstController,
                        decoration: InputDecoration(
                          labelText: 'MNR',
                        ),
                      ),
                    )),
                    Expanded(
                        child: Padding(
                      padding: EdgeInsets.fromLTRB(16.0, 8.0, 16.0, 8.0),
                      child: TextFormField(
                        controller: mnrSecondController,
                        decoration: InputDecoration(
                          labelText: 'MNR',
                        ),
                      ),
                    ))
                  ],
                ),
                Row(
                  children: [
                    Expanded(
                        child: Padding(
                      padding: EdgeInsets.fromLTRB(16.0, 8.0, 16.0, 8.0),
                      child: TextFormField(
                        controller: panFirstController,
                        decoration: InputDecoration(
                          labelText: 'PAN-BOL',
                        ),
                      ),
                    )),
                    Expanded(
                        child: Padding(
                      padding: EdgeInsets.fromLTRB(16.0, 8.0, 16.0, 8.0),
                      child: TextFormField(
                        controller: panSecondController,
                        decoration: InputDecoration(
                          labelText: 'PAN-BOL',
                        ),
                      ),
                    ))
                  ],
                ),
                Row(
                  children: [
                    Expanded(
                        child: Padding(
                      padding: EdgeInsets.fromLTRB(16.0, 8.0, 16.0, 8.0),
                      child: TextFormField(
                        controller: validoFirstController,
                        decoration: InputDecoration(
                          labelText: 'Votos válidos',
                        ),
                      ),
                    )),
                    Expanded(
                        child: Padding(
                      padding: EdgeInsets.fromLTRB(16.0, 8.0, 16.0, 8.0),
                      child: TextFormField(
                        controller: validoSecondController,
                        decoration: InputDecoration(
                          labelText: 'Votos válidos',
                        ),
                      ),
                    ))
                  ],
                ),
                Row(
                  children: [
                    Expanded(
                        child: Padding(
                      padding: EdgeInsets.fromLTRB(16.0, 8.0, 16.0, 8.0),
                      child: TextFormField(
                        controller: blancoFirstController,
                        decoration: InputDecoration(
                          labelText: 'Votos blancos',
                        ),
                      ),
                    )),
                    Expanded(
                        child: Padding(
                      padding: EdgeInsets.fromLTRB(16.0, 8.0, 16.0, 8.0),
                      child: TextFormField(
                        controller: blancoSecondController,
                        decoration: InputDecoration(
                          labelText: 'Votos blancos',
                        ),
                      ),
                    ))
                  ],
                ),
                Row(
                  children: [
                    Expanded(
                        child: Padding(
                      padding: EdgeInsets.fromLTRB(16.0, 8.0, 16.0, 8.0),
                      child: TextFormField(
                        controller: nuloFirstController,
                        decoration: InputDecoration(
                          labelText: 'Votos nulos',
                        ),
                      ),
                    )),
                    Expanded(
                        child: Padding(
                      padding: EdgeInsets.fromLTRB(16.0, 8.0, 16.0, 8.0),
                      child: TextFormField(
                        controller: nuloSecondController,
                        decoration: InputDecoration(
                          labelText: 'Votos nulos',
                        ),
                      ),
                    ))
                  ],
                ),
                Row(
                  children: [
                    Expanded(
                        child: Padding(
                      padding: EdgeInsets.fromLTRB(16.0, 8.0, 16.0, 8.0),
                      child: TextFormField(
                        controller: totalController,
                        decoration: InputDecoration(
                          labelText: 'Votos en total',
                        ),
                      ),
                    )),
                    Expanded(
                        child: Padding(
                            padding: EdgeInsets.fromLTRB(16.0, 16.0, 16.0, 0.0),
                            child: Text(
                              '',
                            ))),
                  ],
                )
              ],
            ),
          )
        ],
      ),
      floatingActionButton: FloatingActionButton(
        backgroundColor: Color(0xFF0e3e3f),
        onPressed: () async {

          showDialog(context: context, builder: (context) => loadingDialog);
          Proceeding data = Proceeding(
              int.parse(idController.text),
              int.parse(numeroController.text),
              int.parse(circunscripcionController.text),
              departamentoController.text,
              provinciaController.text,
              municipioController.text,
              localidadController.text,
              recintoController.text,
              int.parse(ccFirstController.text),
              int.parse(fpvFirstController.text),
              int.parse(mtsFirstController.text),
              int.parse(ucsFirstController.text),
              int.parse(masFirstController.text),
              int.parse(vefFirstController.text),
              int.parse(pdcFirstController.text),
              int.parse(mnrFirstController.text),
              int.parse(panFirstController.text),
              int.parse(validoFirstController.text),
              int.parse(blancoFirstController.text),
              int.parse(nuloFirstController.text),
              int.parse(ccSecondController.text),
              int.parse(fpvSecondController.text),
              int.parse(mtsSecondController.text),
              int.parse(ucsSecondController.text),
              int.parse(masSecondController.text),
              int.parse(vefSecondController.text),
              int.parse(pdcSecondController.text),
              int.parse(mnrSecondController.text),
              int.parse(panSecondController.text),
              int.parse(validoSecondController.text),
              int.parse(blancoSecondController.text),
              int.parse(nuloSecondController.text),
              proceeding.dispositivo,
              proceeding.ubicacion,
              int.parse(totalController.text));
          String json = jsonEncode(data);
          // Update data in backend
          try {
            Dio dio = new Dio();
            Response response = await dio
                .post(ApiURL + "api/v1/sw1/send-confirmation-data", data: json);
            Navigator.pop(context);
            if (response.statusCode == 202) {
              showDialog(context: context, builder: (context) => updateDialog);
            } else {
              showDialog(context: context, builder: (context) => okDialog);
            }
          } on DioError catch (e) {
            print(e);
          }
        },
        child: Icon(Icons.send),
      ),
    );
  }
}

class ResponseArguments {
  final String response;

  ResponseArguments(this.response);
}

// Model necessary for save proceeding data
class Proceeding {
  final int id;
  final int numero;
  final int circunscripcion;
  final String departamento;
  final String provincia;
  final String municipio;
  final String localidad;
  final String recinto;
  final int presidenteCC;
  final int presidenteFPV;
  final int presidenteMTS;
  final int presidenteUCS;
  final int presidenteMAS;
  final int presidenteVeF;
  final int presidentePDC;
  final int presidenteMNR;
  final int presidentePAN;
  final int presidenteValido;
  final int presidenteBlanco;
  final int presidenteNulo;
  final int diputadoCC;
  final int diputadoFPV;
  final int diputadoMTS;
  final int diputadoUCS;
  final int diputadoMAS;
  final int diputadoVeF;
  final int diputadoPDC;
  final int diputadoMNR;
  final int diputadoPAN;
  final int diputadoValido;
  final int diputadoBlanco;
  final int diputadoNulo;
  final String dispositivo;
  final String ubicacion;
  final int votosTotal;

  Proceeding(
      this.id,
      this.numero,
      this.circunscripcion,
      this.departamento,
      this.provincia,
      this.municipio,
      this.localidad,
      this.recinto,
      this.presidenteCC,
      this.presidenteFPV,
      this.presidenteMTS,
      this.presidenteUCS,
      this.presidenteMAS,
      this.presidenteVeF,
      this.presidentePDC,
      this.presidenteMNR,
      this.presidentePAN,
      this.presidenteValido,
      this.presidenteBlanco,
      this.presidenteNulo,
      this.diputadoCC,
      this.diputadoFPV,
      this.diputadoMTS,
      this.diputadoUCS,
      this.diputadoMAS,
      this.diputadoVeF,
      this.diputadoPDC,
      this.diputadoMNR,
      this.diputadoPAN,
      this.diputadoValido,
      this.diputadoBlanco,
      this.diputadoNulo,
      this.dispositivo,
      this.ubicacion,
      this.votosTotal);

  Proceeding.fromJson(Map<String, dynamic> json)
      : id = json['id'],
        numero = json['numero'],
        circunscripcion = json['circunscripcion'],
        departamento = json['departamento'],
        provincia = json['provincia'],
        municipio = json['municipio'],
        localidad = json['localidad'],
        recinto = json['recinto'],
        presidenteCC = json['presidenteCC'],
        presidenteFPV = json['presidenteFPV'],
        presidenteMTS = json['presidenteMTS'],
        presidenteUCS = json['presidenteUCS'],
        presidenteMAS = json['presidenteMAS'],
        presidenteVeF = json['presidenteVeF'],
        presidentePDC = json['presidentePDC'],
        presidenteMNR = json['presidenteMNR'],
        presidentePAN = json['presidentePAN'],
        presidenteValido = json['presidenteValido'],
        presidenteBlanco = json['presidenteBlanco'],
        presidenteNulo = json['presidenteNulo'],
        diputadoCC = json['diputadoCC'],
        diputadoFPV = json['diputadoFPV'],
        diputadoMTS = json['diputadoMTS'],
        diputadoUCS = json['diputadoUCS'],
        diputadoMAS = json['diputadoMAS'],
        diputadoVeF = json['diputadoVeF'],
        diputadoPDC = json['diputadoPDC'],
        diputadoMNR = json['diputadoMNR'],
        diputadoPAN = json['diputadoPAN'],
        diputadoValido = json['diputadoValido'],
        diputadoBlanco = json['diputadoBlanco'],
        diputadoNulo = json['diputadoNulo'],
        dispositivo = json['dispositivo'],
        ubicacion = json['ubicacion'],
        votosTotal = json['votosTotal'];

  Map<String, dynamic> toJson() => {
        'id': id,
        'numero': numero,
        'circunscripcion': circunscripcion,
        'departamento': departamento,
        'provincia': provincia,
        'municipio': municipio,
        'localidad': localidad,
        'recinto': recinto,
        'presidenteCC': presidenteCC,
        'presidenteFPV': presidenteFPV,
        'presidenteMTS': presidenteMTS,
        'presidenteUCS': presidenteUCS,
        'presidenteMAS': presidenteMAS,
        'presidenteVeF': presidenteVeF,
        'presidentePDC': presidentePDC,
        'presidenteMNR': presidenteMNR,
        'presidentePAN': presidentePAN,
        'presidenteValido': presidenteValido,
        'presidenteBlanco': presidenteBlanco,
        'presidenteNulo': presidenteNulo,
        'diputadoCC': diputadoCC,
        'diputadoFPV': diputadoFPV,
        'diputadoMTS': diputadoMTS,
        'diputadoUCS': diputadoUCS,
        'diputadoMAS': diputadoMAS,
        'diputadoVeF': diputadoVeF,
        'diputadoPDC': diputadoPDC,
        'diputadoMNR': diputadoMNR,
        'diputadoPAN': diputadoPAN,
        'diputadoValido': diputadoValido,
        'diputadoBlanco': diputadoBlanco,
        'diputadoNulo': diputadoNulo,
        'dispositivo': dispositivo,
        'ubicacion': ubicacion,
        'votosTotal': votosTotal,
      };
}
