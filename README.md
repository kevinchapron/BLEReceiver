# BLE Receiver

[![Languages](https://img.shields.io/badge/languages-En-green.svg)]()
[![Licence Apache2](https://img.shields.io/hexpm/l/plug.svg)](http://www.apache.org/licenses/LICENSE-2.0)

---

This project is a **GoLang application** designed to be used on a RaspBerry Pi3.
It connects through **BLE** (**B**luetooth **L**ow **E**nergy) to a given peripheral.
It suscribes to characteristics available from Peripheral systems that stream theirs own data.

9 Characteristics are available from peripheral, one for each axis of each device. (9-DOF)
It's a good way to make real-times applications through bluetooth devices

This project runs on **[Paypal/Gatt](https://github.com/paypal/gatt/)**.

More information
---

This project has been created for my research in the **[LIARA](http://liara.uqac.ca/)** lab 
(Laboratoire d'Intelligence Ambiante pour la Reconnaissance d'Activités), at the 
« Université du Québec À Chicoutimi (**[UQAC](http://www.uqac.ca/)**) »

Author
---
**[Kévin CHAPRON](http://kevin-chapron.fr/)** - _2016_

Download
---
To download and use this project, just use the command below : 
> ```go get github.com/kevinchapron/BLEReceiver```

To download sources through git, just clone it using the command below : 
> ```git clone https://github.com/kevinchapron/BLEReceiver.git```

License
---
    Copyright 2016 Kévin Chapron

    Licensed under the Apache License, Version 2.0 (the "License");
    you may not use this file except in compliance with the License.
    You may obtain a copy of the License at

        http://www.apache.org/licenses/LICENSE-2.0

    Unless required by applicable law or agreed to in writing, software
    distributed under the License is distributed on an "AS IS" BASIS,
    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
    See the License for the specific language governing permissions and
    limitations under the License.
