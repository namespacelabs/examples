from django.urls import include, path
from django.views.generic import TemplateView
from .views import home, create_item, list_items

urlpatterns = [
    path("", home, name="item-home"),
    path("create/", create_item, name="item-create"),
    path("list/", list_items, name="item-list"),
]
