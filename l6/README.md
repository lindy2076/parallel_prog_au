# Лабораторная работа №6 «Производители–Потребители»

В научно-технической литературе существует достаточно большое количество
вариантов постановки данной задачи. В наиболее простом случае предполагается,
что существует два потока, один из которых (производитель) генерирует
сообщения (изделия), а второй поток (потребитель) их принимает для
последующей обработки. Потоки взаимодействуют через некоторую область
памяти (хранилище), в которой производитель размещает свои генерируемые
сообщения и из которой эти сообщения извлекаются потребителем

производитель -> хранилище сообщений -> потребитель

Рассмотрев постановку данной задачи, можно заметить, что хранилище сообщений
представляет собой не что иное, как общий разделяемый ресурс, и использование
этого ресурса должно быть построено по правилам взаимоисключения. Кроме того,
следует учитывать, что потребление ресурса иногда может оказаться
невозможным (отсутствие сообщений в хранилище), а при добавлении сообщений
в хранилище могут происходить задержки (в случае полного заполнения
хранилища).
