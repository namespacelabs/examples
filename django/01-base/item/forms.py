from django import forms


class ItemCreateForm(forms.Form):
    name = forms.CharField(max_length=50)
