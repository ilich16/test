import 'dart:io';
import 'package:uuid/uuid.dart';
import 'package:flutter/material.dart';
import 'package:image_picker/image_picker.dart';
import 'package:dio/dio.dart';
import 'package:path/path.dart' show join;
import 'package:second_test/screens/proceeding_screen.dart';
import 'package:location/location.dart';
import 'package:device_info/device_info.dart';

// The backend's URL
const ApiURL = 'https://example.com:8080/';

// The picker image screen for mobile app
class PickerImageScreen extends StatefulWidget {
  @override
  _PickerImageScreenState createState() => _PickerImageScreenState();
}

class _PickerImageScreenState extends State<PickerImageScreen> {
  // Variables for open the file
  File _image;
  final picker = ImagePicker();
  String imagePath;
  // Variable for save location
  Location location = new Location();

  // Function for open the camera
  Future getImageCamera() async {
    final pickedFile = await picker.getImage(source: ImageSource.camera);

    setState(() {
      _image = File(pickedFile.path);
      imagePath = pickedFile.path;
    });
  }

  // Function for open the file system
  Future getImageFile() async {
    final pickedFile = await picker.getImage(source: ImageSource.gallery);

    setState(() {
      _image = File(pickedFile.path);
      imagePath = pickedFile.path;
    });
  }

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

  // Function for send data to backend
  void sendImageToServer(String filepath) async {
    bool _serviceEnabled;
    PermissionStatus _permissionGranted;
    LocationData _locationData;

    // Verify if the app has the required permissions
    _serviceEnabled = await location.serviceEnabled();
    if (!_serviceEnabled) {
      _serviceEnabled = await location.requestService();
      if (!_serviceEnabled) {
        return;
      }
    }

    _permissionGranted = await location.hasPermission();
    if (_permissionGranted == PermissionStatus.denied) {
      _permissionGranted = await location.requestPermission();
      if (_permissionGranted != PermissionStatus.granted) {
        return;
      }
    }

    // Get location of mobile device
    _locationData = await location.getLocation();
    String ubicacion = _locationData.latitude.toString() + ',' + _locationData.longitude.toString();

    // Get information of mobile device
    DeviceInfoPlugin deviceInfo = DeviceInfoPlugin();
    AndroidDeviceInfo androidInfo = await deviceInfo.androidInfo;

    // Show loading widget
    showDialog(context: context, builder: (context) => loadingDialog);
    var uuid = Uuid();
    FormData formData = new FormData.fromMap({
      "image":
          await MultipartFile.fromFile(filepath, filename: '${uuid.v1()}.jpg'),
      "location":
          ubicacion,
      "device":
          androidInfo.model,
    });
    // Send data to backend
    try {
      Response response =
          await Dio().post(ApiURL + "api/v1/sw1/send-image", data: formData);
      String message = response.toString();
      Navigator.of(context).pushReplacementNamed('/proceeding', arguments: ResponseArguments(message));
    } catch (e) {
      print(e);
    }
  }

  @override
  Widget build(BuildContext context) {
    final AlertDialog sourceDialog = AlertDialog(
      title: Text(
        'Selecionar fuente',
        textAlign: TextAlign.center,
      ),
      content: Row(
        mainAxisAlignment: MainAxisAlignment.spaceBetween,
        children: [
          OutlineButton(
            color: Color(0xFF0e3e3f),
            onPressed: () {
              Navigator.pop(context);
              getImageCamera();
            },
            child: Text(
              "CÁMARA",
              style: TextStyle(color: Color(0xFF0e3e3f)),
            ),
          ),
          OutlineButton(
            color: Color(0xFF0e3e3f),
            onPressed: () {
              Navigator.pop(context);
              getImageFile();
            },
            child: Text("GALERÍA", style: TextStyle(color: Color(0xFF0e3e3f))),
          )
        ],
      ),
    );

    return Scaffold(
      appBar: AppBar(
        title: Text('SELECCIONAR IMAGEN'),
      ),
      body: Center(
        child: Column(
          children: [
            _image == null
                ? Padding(
                    padding: EdgeInsets.all(16),
                    child: Text(
                      'Por favor, selecciona una imagen presionando el botón que se encuentra en la parte inferior derecha de tu pantalla.',
                      textAlign: TextAlign.center,
                      style: TextStyle(color: Color(0xFF0e3e3f)),
                    ),
                  )
                : Image.file(_image),
            _image == null
                ? Text("")
                : RaisedButton(
                    color: Color(0xFF0e3e3f),
                    textColor: Colors.white,
                    child: Text('ENVIAR IMAGEN',
                        style: TextStyle(
                            fontSize: 14, fontWeight: FontWeight.w800)),
                    onPressed: () {
                      sendImageToServer(imagePath);
                    },
                  ),
          ],
        ),
      ),
      floatingActionButton: FloatingActionButton(
        backgroundColor: Color(0xFF0e3e3f),
        //onPressed: getImage,
        onPressed: () {
          showDialog(context: context, builder: (context) => sourceDialog);
        },
        tooltip: 'Pick Image',
        child: Icon(Icons.add_a_photo),
      ),
    );
  }
}
