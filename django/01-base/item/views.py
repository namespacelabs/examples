from django.shortcuts import render
from django.http import HttpResponseNotAllowed
from .models import Item
from .forms import ItemCreateForm


def home(request):
    form = ItemCreateForm()
    return render(
        request, "index.html", context={"items": Item.objects.all(), "form": form}
    )


def create_item(request):
    if request.method == "POST":
        form = ItemCreateForm(request.POST)
        if form.is_valid():
            Item.objects.create(name=form.cleaned_data["name"])
            response = render(
                request,
                "form.html",
                context={"form": ItemCreateForm()},
            )
            response["HX-Trigger"] = "itemCreated"
            return response
        else:
            return render(request, "error.html", context={"form": form})
    return HttpResponseNotAllowed(permitted_methods=["POST"])


def list_items(request):
    return render(request, "list.html", context={"items": Item.objects.all()})
