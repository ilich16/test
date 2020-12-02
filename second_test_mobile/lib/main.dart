import 'package:flutter/material.dart';
import 'package:camera/camera.dart';
import 'package:second_test/screens/login_screen.dart';
import 'package:second_test/screens/home_screen.dart';
import 'package:second_test/screens/pick_image_screen.dart';
import 'package:second_test/screens/proceeding_screen.dart';

Future<void> main() async {
  // Camera controller
  // Ensure that plugin services are initialized so that 'availableCameras()'
  // can be called before 'runApp()'
  WidgetsFlutterBinding.ensureInitialized();

  // Obtain a list of the available cameras on the device.
  final cameras = await availableCameras();
  // Get a specific camera from the list of available cameras.
  final camera = cameras.first;

  runApp(MaterialApp(
    title: '2do Parcial',
    theme: ThemeData(
      primaryColor: Color(0xFF0e3e3f),
      fontFamily: 'Open Sans',
    ),
    initialRoute: '/login',
    routes: {
      '/login': (context) => LoginScreen(),
      '/home': (context) => HomeScreen(),
      '/camera': (context) => PickerImageScreen(),
      '/proceeding': (context) => ProceedingScreen()
    },
    debugShowCheckedModeBanner: false,
  ));
}
