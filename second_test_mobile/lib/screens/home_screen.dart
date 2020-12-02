import 'package:flutter/material.dart';

// The home screen for the mobile app
class HomeScreen extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text('PÁGINA PRINCIPAL'),
      ),
      body: Column(
        children: [
          Padding(
            padding: EdgeInsets.all(16),
            child: Image.asset('assets/secondtest.png'),
          ),
          Padding(
            padding: EdgeInsets.all(16),
            child: Text(
              'Lleva el conteo de votos de una manera mas rápida y eficiente, logrando así que las elecciones de este año se realicen de una manera más limpia y transparente.',
              textAlign: TextAlign.center,
              style: TextStyle(color: Color(0xFF0e3e3f)),
            ),
          ),
          RaisedButton(
            color: Color(0xFF0e3e3f),
            textColor: Colors.white,
            child: Text('EMPEZAR',
                style: TextStyle(fontSize: 14, fontWeight: FontWeight.w800)),
            onPressed: () {
              // Navigate to the camera screen using a named route.
              Navigator.pushNamed(context, '/camera');
            },
          )
        ],
      ),
    );
  }
}
