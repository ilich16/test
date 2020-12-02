import 'package:flutter/material.dart';
import 'package:dio/dio.dart';

// The backend's URL
const ApiURL = 'https://example.com:8080/';

// The login screen for the mobile app
class LoginForm extends StatefulWidget {
  @override
  _LoginFormState createState() => _LoginFormState();
}

class LoginScreen extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        // Here we take the value from the MyHomePage object that was created by
        // the App.build method, and use it to set our appbar title.
        title: Text("INICIAR SESIÓN"),
      ),
      body: LoginForm(),
    );
  }
}

// Create a corresponding State class.
// This class holds data related to the form.
class _LoginFormState extends State<LoginForm> {
  // Create a global key that uniquely identifies the Form widget
  // and allows validation of the form.
  final _formKey = GlobalKey<FormState>();
  // Controllers for the TextFormFields
  final usernameController = TextEditingController();
  final passwordController = TextEditingController();
  // Switch for show password
  bool _passwordVisible;

  @override
  void initState() {
    _passwordVisible = false;
  }

  @override
  void dispose() {
    // Clean up the controllers when the widget is disposed.
    usernameController.dispose();
    passwordController.dispose();
    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    final AlertDialog errorDialog = AlertDialog(
      title: Text('Error'),
      content: Text('Por favor, verifique que sus datos sean correctos.'),
      actions: [
        FlatButton(
          onPressed: () => Navigator.pop(context),
          child: Text('REGRESAR'),
        )
      ],
    );

    // Build a Form widget using the _formKey created above.
    return Form(
        key: _formKey,
        child: Column(
          crossAxisAlignment: CrossAxisAlignment.start,
          children: [
            Padding(
              padding: EdgeInsets.all(16.0),
              child: TextFormField(
                controller: usernameController,
                decoration: InputDecoration(
                  icon: Icon(Icons.person),
                  border: OutlineInputBorder(),
                  labelText: 'Nombre de usuario',
                ),
                validator: (value) {
                  if (value.isEmpty) {
                    return 'Por favor, ingrese su nombre de usuario';
                  }
                  return !value.contains('@')
                      ? 'Por favor, ingrese un correo válido'
                      : null;
                },
              ),
            ),
            Padding(
              padding: EdgeInsets.fromLTRB(16, 0, 16, 8),
              child: TextFormField(
                controller: passwordController,
                decoration: InputDecoration(
                    icon: Icon(Icons.vpn_key),
                    border: OutlineInputBorder(),
                    labelText: 'Contraseña',
                    suffixIcon: IconButton(
                      icon: Icon(_passwordVisible
                          ? Icons.visibility
                          : Icons.visibility_off),
                      onPressed: () {
                        setState(() {
                          _passwordVisible = !_passwordVisible;
                        });
                      },
                    )),
                obscureText: !_passwordVisible,
                validator: (value) {
                  if (value.isEmpty) {
                    return 'Por favor, ingrese su contraseña';
                  }
                  return null;
                },
              ),
            ),
            Padding(
              padding: EdgeInsets.symmetric(vertical: 16.0),
              child: Center(
                child: RaisedButton(
                  color: Color(0xFF0e3e3f),
                  textColor: Colors.white,
                  onPressed: () async {
                    // Validate returns true if the form is valid, or false
                    // otherwise.
                    if (_formKey.currentState.validate()) {
                      String username = usernameController.text;
                      String password = passwordController.text;
                      try {
                        Dio dio = new Dio();
                        await dio.post(ApiURL + "api/v1/sw1/login", data: {
                          "username": "$username",
                          "password": "$password"
                        });
                        // Navigate to the home screen using a named route.
                        Navigator.pushNamed(context, '/home');
                      } on DioError catch (e) {
                        showDialog(
                            context: context, builder: (context) => errorDialog);
                      }
                    }
                  },
                  child: Text('INICIAR SESIÓN', style: TextStyle(fontSize: 14, fontWeight: FontWeight.w800)),
                ),
              )
            )
          ],
        ));
  }
}
