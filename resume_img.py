from pdf2image import convert_from_path

for index, img in enumerate(convert_from_path('images/Shreyas Donti Resume.pdf')):
    img.save('images/Shreyas Donti Resume.png', 'PNG')

 